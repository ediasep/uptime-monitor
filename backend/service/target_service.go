package service

import (
	"uptime-monitor/model"
	"uptime-monitor/repository"
)

type TargetService struct {
	repo *repository.TargetRepository
}

func NewTargetService(repo *repository.TargetRepository) *TargetService {
	return &TargetService{repo: repo}
}

func (s *TargetService) CreateTarget(name, url string, interval int) (model.Target, error) {
	return s.repo.Add(name, url, interval)
}

func (s *TargetService) GetAllTargets() ([]model.Target, error) {
	return s.repo.List()
}
