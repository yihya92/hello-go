package userstore

import (
	"errors"
	"sync"
)

type User struct {
	ID   string
	Name string
	Age  int
}

var (
	ErrUserExists   = errors.New("User already Exists")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidID    = errors.New("invalid user id")
)

type Store struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewStore() *Store {
	return &Store{
		users: make(map[string]User),
	}
}

// Add inserts a new user.
// Returns ErrInvalidID if ID is empty.
// Returns ErrUserExists if duplicate.
func (s *Store) Add(u User) error {
	if u.ID == "" {
		return ErrInvalidID
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[u.ID]; exists {
		return ErrUserExists
	}

	s.users[u.ID] = u
	return nil
}

// Get retrieves a user by ID.
func (s *Store) Get(id string) (User, error) {
	if id == "" {
		return User{}, ErrInvalidID
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return u, nil
}

// List returns a snapshot slice of all users.
func (s *Store) List() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}

	return users
}

// Delete removes a user by ID.
func (s *Store) Delete(id string) error {
	if id == "" {
		return ErrInvalidID
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return ErrUserNotFound
	}

	delete(s.users, id)
	return nil
}
