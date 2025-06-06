package service

import (
	"uptime-monitor/model"
	target "uptime-monitor/repository/target"
)

type TargetService struct {
	repo target.TargetRepository // now using interface
}

func NewTargetService(repo target.TargetRepository) *TargetService {
	return &TargetService{repo: repo}
}

func (s *TargetService) CreateTarget(name, url string, interval int) (model.Target, error) {
	return s.repo.Add(name, url, interval)
}

func (s *TargetService) GetAllTargets() ([]model.Target, error) {
	return s.repo.List()
}
