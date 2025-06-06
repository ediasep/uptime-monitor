package user

import (
	"errors"
	"sync"
	"uptime-monitor/model"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	GetByID(id int64) (*model.User, error)
	Create(user *model.User) error
}

type userRepository struct {
	mu    sync.RWMutex
	users map[int64]*model.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: make(map[int64]*model.User),
	}
}

func (r *userRepository) GetByID(id int64) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *userRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}
