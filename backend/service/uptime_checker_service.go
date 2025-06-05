package service

import (
	"log"
	"net/http"
	"time"
	"uptime-monitor/model"
	"uptime-monitor/repository"
)

type UptimeCheckerService struct {
	targetRepo *repository.TargetRepository
	logRepo    *repository.TargetLogRepository
	eventSvc   *EventService
	interval   time.Duration
	stopChan   chan struct{}
}

func NewUptimeCheckerService(
	targetRepo *repository.TargetRepository,
	logRepo *repository.TargetLogRepository,
	eventSvc *EventService,
	interval time.Duration,
) *UptimeCheckerService {
	return &UptimeCheckerService{
		targetRepo: targetRepo,
		logRepo:    logRepo,
		eventSvc:   eventSvc,
		interval:   interval,
		stopChan:   make(chan struct{}),
	}
}

func (u *UptimeCheckerService) Start() {
	go func() {
		ticker := time.NewTicker(u.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				u.checkAllTargets()
			case <-u.stopChan:
				log.Println("Uptime checker stopped.")
				return
			}
		}
	}()
}

func (u *UptimeCheckerService) Stop() {
	close(u.stopChan)
}

func (u *UptimeCheckerService) checkAllTargets() {
	targets, err := u.targetRepo.List()
	if err != nil {
		log.Println("Failed to list targets:", err)
		return
	}

	for _, target := range targets {
		go u.checkTarget(target)
	}
}

func (u *UptimeCheckerService) checkTarget(target model.Target) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(target.URL)

	status := "DOWN"
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 400 {
		status = "UP"
	}

	log.Printf("[UptimeCheck] %s (%s) is %s", target.Name, target.URL, status)

	// Save status log
	saveErr := u.logRepo.Save(model.NewTargetLog(target.ID, status))
	if saveErr != nil {
		log.Println("Failed to save log:", saveErr)
		return
	}

	// Check for 3 consecutive failures
	if status == "DOWN" {
		failCount, err := u.logRepo.CountRecentFailures(target.ID, 3)
		if err == nil && failCount >= 3 {
			u.eventSvc.HandleTargetDown(target.ID)
		}
	}
}
