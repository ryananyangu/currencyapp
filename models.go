package main

import "time"

// Currency model for storing currency information
type Currency struct {
	Country     string
	Name        string
	LastFetchAt time.Time
}
