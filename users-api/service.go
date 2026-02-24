package main

// business logic layer

import (
	"errors"
	"sync"
)

type UserService struct {
	users  map[int]User
	nextID int        // auto increment for ids
	mu     sync.Mutex // Only one goroutine modifies the map at a time
}

// constructor function
// userservice as pointer in oder not to copy mutex
func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *UserService) GetAll() []User {
	s.mu.Lock()
	defer s.mu.Unlock()

	list := []User{}
	for _, user := range s.users {
		list = append(list, user)
	}
	return list
}

func (s *UserService) GetByID(id int) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) Create(user User) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user

	return user
}

func (s *UserService) Update(id int, updated User) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	updated.ID = id
	s.users[id] = updated

	return updated, nil
}

func (s *UserService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}
