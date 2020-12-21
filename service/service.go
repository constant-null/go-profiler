package service

import "time"

// Service contains information about service
type Service struct {
	ID        int64
	Name      string
	UpdatedAt time.Time
}
