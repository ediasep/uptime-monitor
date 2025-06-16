package model

import (
	"time"

	"github.com/google/uuid"
)

// DailyUptimeResponse represents the daily uptime percentage response
type DailyUptimeResponse struct {
	TargetID         uuid.UUID `json:"target_id"`
	Date             time.Time `json:"date"`
	UptimePercentage float64   `json:"uptime_percentage"`
}

// NewDailyUptimeResponse creates a new DailyUptimeResponse instance
func NewDailyUptimeResponse(targetID uuid.UUID, date time.Time, uptimePercentage float64) *DailyUptimeResponse {
	return &DailyUptimeResponse{
		TargetID:         targetID,
		Date:             date,
		UptimePercentage: uptimePercentage,
	}
}
