package service

import (
	"uptime-monitor/model"
	"uptime-monitor/repository/targetlog"
)

type TargetLogService struct {
	repo targetlog.TargetLogRepository
}

func NewTargetLogService(repo targetlog.TargetLogRepository) *TargetLogService {
	return &TargetLogService{repo: repo}
}

func (s *TargetLogService) CreateLog(log model.TargetLog) error {
	return s.repo.Save(log)
}

func (s *TargetLogService) GetLogsByTargetID(targetID string) ([]model.TargetLog, error) {
	return s.repo.ListByTargetID(targetID)
}

func (s *TargetLogService) DeleteLogsByTargetID(targetID string) error {
	return s.repo.DeleteByTargetID(targetID)
}

func (s *TargetLogService) GetDailyUptimePercentageByTargetID(targetID string) ([]model.DailyUptimeResponse, error) {
	dailyUptime, err := s.repo.GetDailyUptimePercentageByTargetID(targetID)
	if err != nil {
		return nil, err
	}
	return dailyUptime, nil
}
