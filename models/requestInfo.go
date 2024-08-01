package models

import (
	"time"

	"github.com/google/uuid"
)

// This is the info stored into
type RequestInfo struct {
	Port      string    `json:"port"`
	Timestamp time.Time `json:"timestamp"`
	IPAddress string    `json:"ip_address"`
	Path      string    `json:"path"`
}

func (r *RequestInfo) ConvertToDBRecord() *DBRecord {
	return &DBRecord{
		ID:        uuid.New().String(),
		Port:      r.Port,
		IPAddress: r.IPAddress,
		Timestamp: r.Timestamp,
		Path:      r.Path,
	}
}
