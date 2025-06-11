package model

import (
	"time"

	"github.com/google/uuid"
)

type TargetListDto struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	URL           string     `json:"url"`
	Interval      int        `json:"interval"`
	CreatedAt     time.Time  `json:"created_at"`
	LastCheckedAt *time.Time `json:"last_checked_at"`
	LastStatus    *string    `json:"last_status"`
}

func NewTargetListDto(name, url string, interval int, lastCheckedAt *time.Time, status *string) TargetListDto {
	return TargetListDto{
		ID:            uuid.NewString(),
		Name:          name,
		URL:           url,
		Interval:      interval,
		CreatedAt:     time.Now(),
		LastCheckedAt: lastCheckedAt, // Initialize to zero value
		LastStatus:    status,        // Default status
	}
}
