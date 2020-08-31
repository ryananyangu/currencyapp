package main

import "time"

// Currency model for storing currency information
type Currency struct {
	Country     string
	Name        string
	LastFetchAt time.Time
	Code        string
}

// Language model for storing available retieved languages
type Language struct {
	Name        string
	Code        string
	LastFetchAt time.Time
}
