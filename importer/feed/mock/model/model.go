package model

// APIResponse describes the mock feed's response
type APIResponse struct {
	Horses []Horse `json:"horses"`
	Odds   []Odds  `json:"odds"`
}

// Horse represents a horse's information
type Horse struct {
	Name   string `json:"name"`
	LineID int    `json:"lineID"`
}

// Odds represents an object in the odds array
type Odds struct {
	ID  int     `json:"id"`
	Odd float64 `json:"odds"`
}
