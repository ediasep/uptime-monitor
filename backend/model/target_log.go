package model

import (
	"time"

	"github.com/google/uuid"
)

type TargetLog struct {
	ID        string
	TargetID  string
	Status    string
	Timestamp time.Time
}

func NewTargetLog(targetID, status string) TargetLog {
	return TargetLog{
		ID:        uuid.NewString(),
		TargetID:  targetID,
		Status:    status,
		Timestamp: time.Now(),
	}
}
