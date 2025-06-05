package service

import "log"

type EventService struct{}

func NewEventService() *EventService {
	return &EventService{}
}

func (e *EventService) HandleTargetDown(targetID string) {
	// TODO: implement alert/notification
	log.Printf("!!! ALERT: Target ID %s is DOWN 3 times in a row !!!", targetID)
}
