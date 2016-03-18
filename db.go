package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/Clever/optimus.v3/sources/csv"
)

// LatLon contains a latitude and a longitude
type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (ll *LatLon) String() string {
	return fmt.Sprintf("%f,%f", ll.Lat, ll.Lon)
}

// GeoDB is a geo database that enables translating IP addresses to locations
type GeoDB struct {
	blockBoundaries    []int
	boundaryToLocation map[int]*LatLon
}

func piecesToInt(pieces []string) (int, error) {
	parts := [4]int{}

	var piece int
	var err error
	piece, err = convertToInt(pieces[0])
	if err != nil {
		return 0, err
	}
	parts[0] = piece
	piece, err = convertToInt(pieces[1])
	if err != nil {
		return 0, err
	}
	parts[1] = piece
	piece, err = convertToInt(pieces[2])
	if err != nil {
		return 0, err
	}
	parts[2] = piece
	piece, err = convertToInt(pieces[3])
	if err != nil {
		return 0, err
	}
	parts[3] = piece

	return parts[0]*256*256*256 + parts[1]*256*256 + parts[2]*256 + parts[3], nil
}

// Lookup takes in an IPv4 IP address and returns a LatLon that that IP corresponds to.
func (g *GeoDB) Lookup(ip string) (*LatLon, error) {
	pieces := strings.Split(ip, ".")
	if len(pieces) != 4 {
		return nil, fmt.Errorf("invalid IP %s", ip)
	}

	num, err := piecesToInt(pieces)
	if err != nil {
		return nil, fmt.Errorf("invalid IP pieces %s", ip)
	}

	idx := sort.SearchInts(g.blockBoundaries, num)

	// There are three possible scenarios:
	// 1. The number matches a start boundary exactly
	// 2. The number matches an end boundary exactly
	// 3. The number is between a start and an end boundary
	// We'll check the idx +/- 1 to see if it matches any exactly, if so, we use that index.
	// Otherwise we use the idx-1, which should be the start boundary
	destIdx := idx - 1
	if num == g.blockBoundaries[idx] {
		destIdx = idx
	}
	if num == g.blockBoundaries[idx+1] {
		destIdx = idx + 1
	}

	return g.boundaryToLocation[g.blockBoundaries[destIdx]], nil
}

// NewGeoDB constructs a new GeoDB using the location and block data at the given path.
func NewGeoDB(path string) (*GeoDB, error) {
	db := &GeoDB{
		blockBoundaries:    []int{},
		boundaryToLocation: map[int]*LatLon{},
	}

	locations, err := os.Open(path + "/GeoLiteCity-Location.csv")
	if err != nil {
		return nil, err
	}
	defer locations.Close()

	locToLatLon := map[int]*LatLon{}

	t := csv.New(locations)
	for record := range t.Rows() {
		loc, err := convertToInt(record["locId"])
		if err != nil {
			return nil, err
		}
		lat, err := convertToFloat(record["latitude"])
		if err != nil {
			return nil, err
		}
		lon, err := convertToFloat(record["longitude"])
		if err != nil {
			return nil, err
		}

		locToLatLon[loc] = &LatLon{Lat: lat, Lon: lon}
	}
	if t.Err() != nil {
		return nil, t.Err()
	}

	blocks, err := os.Open(path + "/GeoLiteCity-Blocks.csv")
	if err != nil {
		return nil, err
	}
	defer blocks.Close()

	t = csv.New(blocks)
	for record := range t.Rows() {
		start, err := convertToInt(record["startIpNum"])
		if err != nil {
			return nil, fmt.Errorf("failed to find start: %s", err)
		}
		end, err := convertToInt(record["endIpNum"])
		if err != nil {
			return nil, fmt.Errorf("failed to find end: %s", err)
		}
		loc, err := convertToInt(record["locId"])
		if err != nil {
			return nil, fmt.Errorf("failed to find loc: %s", err)
		}

		db.blockBoundaries = append(db.blockBoundaries, start)
		db.blockBoundaries = append(db.blockBoundaries, end)
		db.boundaryToLocation[start] = locToLatLon[loc]
		db.boundaryToLocation[end] = locToLatLon[loc]
	}
	if t.Err() != nil {
		return nil, t.Err()
	}

	sort.Ints(db.blockBoundaries)

	return db, nil
}

func convertToInt(field interface{}) (int, error) {
	str := field.(string)
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("couldn't convert string %s", str)
	}
	return i, nil
}

func convertToFloat(field interface{}) (float64, error) {
	str := field.(string)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, fmt.Errorf("couldn't convert string %s", str)
	}
	return f, nil
}
