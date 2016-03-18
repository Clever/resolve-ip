package main

import "testing"

func TestNewDB(t *testing.T) {
	_, err := NewGeoDB(*path)
	if err != nil {
		t.Fatal(err)
	}
}
