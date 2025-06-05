package service

import (
	"uptime-monitor/model"
	"uptime-monitor/repository"
)

type UserService interface {
	GetUser(id int64) (*model.User, error)
	CreateUser(user *model.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id int64) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}
