package userstore

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidAge = errors.New("invalid age")
	ErrEmptyName  = errors.New("name is required")
)

// The service depends on an interface
type Service struct {
	repo UserRepository
}

// Constructor
func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

// This is business logic for creating users.
func (s *Service) CreateUser(u User) error {
	if u.ID == "" {
		return ErrInvalidID
	}
	if u.Name == "" {
		return ErrEmptyName
	}
	if u.Age < 0 {
		return ErrInvalidAge
	}
	// Call Repository
	if err := s.repo.Add(u); err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	//If everything passes, no error.
	return nil
}

func (s *Service) GetUser(id string) (User, error) {
	u, err := s.repo.Get(id)
	if err != nil {
		return User{}, fmt.Errorf("get user: %w", err)
	}
	return u, nil
}
