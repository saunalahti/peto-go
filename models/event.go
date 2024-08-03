package models

type Event struct {
	ID          string `json:"id"`
	Location    string `json:"location"`
	LocationInt string `json:"location_int,omitempty"`
	Incident    string `json:"incident"`
	Severity    int    `json:"severity,omitempty"`
	Datetime    string `json:"date"`
}
