package service

import (
	"errors"
	"uptime-monitor/model"
	target "uptime-monitor/repository/target"
)

var ErrTargetNotFound = errors.New("target not found")

type TargetService struct {
	repo target.TargetRepository
}

func NewTargetService(repo target.TargetRepository) *TargetService {
	return &TargetService{repo: repo}
}

func (s *TargetService) CreateTarget(name, url string, interval int) (model.Target, error) {
	return s.repo.Add(name, url, interval)
}

func (s *TargetService) UpdateTarget(id, name, url string, interval int) (model.Target, error) {
	return s.repo.Update(id, name, url, interval)
}

func (s *TargetService) GetAllTargets() ([]model.TargetListDto, error) {
	return s.repo.List()
}

func (s *TargetService) GetTargetByID(id string) (model.Target, error) {
	return s.repo.GetByID(id)
}

func (s *TargetService) DeleteTarget(id string) error {
	return s.repo.Delete(id)
}
