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

// ConvResponse response struct from conversion api
type ConvResponse struct {
	success bool          `json:"success"`
	terms   string        `json:"terms"`
	privacy string        `json:"privacy"`
	query   ResponseQuery `json:"query"`
	info    ResponseInfo  `json:"info"`
	result  string        `json:"result"`
}

// ResponseQuery substructure of Convertion response
type ResponseQuery struct {
	from   string `json:"from"`
	to     string `json:"to"`
	amount int    `json:"amount"`
}

// ResponseInfo substructure of conversion response
type ResponseInfo struct {
	timestamp string  `json:"timestamp"`
	quote     float32 `json:"quote"`
}
