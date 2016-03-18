package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db *GeoDB

func init() {
	// Just to make benchmarks faster
	var err error
	db, err = NewGeoDB(*path)
	if err != nil {
		log.Fatal(err)
	}
}

func TestDB(t *testing.T) {
	assert := assert.New(t)

	latlon, err := db.Lookup("216.75.229.114")
	assert.Nil(err)

	assert.Equal(latlon, &LatLon{37.555, -122.2687})
}

func BenchmarkLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db.Lookup("216.75.229.114")
	}
}
