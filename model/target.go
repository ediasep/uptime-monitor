package model

import (
	"time"

	"github.com/google/uuid"
)

type Target struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Interval  int       `json:"interval"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTarget(name, url string, interval int) Target {
	return Target{
		ID:        uuid.NewString(),
		Name:      name,
		URL:       url,
		Interval:  interval,
		CreatedAt: time.Now(),
	}
}
