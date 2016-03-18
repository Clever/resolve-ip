package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/Clever/optimus.v3/sources/csv"
)

// ErrIPMissing is return when the input IP address doesn't exist in any of our IP blocks
var ErrIPMissing = errors.New("IP address isn't in any of our IP blocks")

// LatLon contains a latitude and a longitude
type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (ll *LatLon) String() string {
	return fmt.Sprintf("%f,%f", ll.Lat, ll.Lon)
}

// Boundary contains the start and end integers of an IP block and the LatLon that it corresponds to
type Boundary struct {
	Start, End int
	*LatLon
}

type byStart []*Boundary

func (a byStart) Len() int           { return len(a) }
func (a byStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

// GeoDB is a geo database that enables translating IP addresses to locations
type GeoDB struct {
	boundaries []*Boundary
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

	// sort.Search returns the smallest_ index for which the function returns true. Given that, we
	// just need to return whether or not our IP is less than the end of the boundary.
	idx := sort.Search(len(g.boundaries), func(idx int) bool {
		boundary := g.boundaries[idx]
		return num < boundary.End
	})
	// It's possible that our IP address isn't in a boundary, in which case it will return the index
	// of the first boundary greater than our IP address. We can check for this case by making sure
	// that the boundary start is smaller than our number.
	if g.boundaries[idx].Start > num {
		return nil, ErrIPMissing
	}

	return g.boundaries[idx].LatLon, nil
}

// NewGeoDB constructs a new GeoDB using the location and block data at the given path.
func NewGeoDB(path string) (*GeoDB, error) {
	db := &GeoDB{
		boundaries: []*Boundary{},
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

		boundary := &Boundary{Start: start, End: end, LatLon: locToLatLon[loc]}
		db.boundaries = append(db.boundaries, boundary)
	}
	if t.Err() != nil {
		return nil, t.Err()
	}

	sort.Sort(byStart(db.boundaries))

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
