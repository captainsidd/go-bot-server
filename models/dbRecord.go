package models

import "time"

// This is the info stored into
type DBRecord struct {
	ID          string    `json:"id"`
	Port        string    `json:"port"`
	Path        string    `json:"path"`
	Timestamp   time.Time `json:"timestamp"`
	CreatedAt   time.Time `json:"created_at"`
	IPAddress   string    `json:"ip_address"`
	CountryCode string    `json:"country_code"`
	City        string    `json:"city"`
}
