package userstore

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	store := NewStore()

	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{"valid user", User{ID: "1", Name: "Alice"}, nil},
		{"duplicate user", User{ID: "1", Name: "Alice"}, ErrUserExists},
		{"empty id", User{Name: "Bob"}, ErrInvalidID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Add(tt.user)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	store := NewStore()
	_ = store.Add(User{ID: "1", Name: "Alice"})

	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{"existing user", "1", nil},
		{"missing user", "999", ErrUserNotFound},
		{"empty id", "", ErrInvalidID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := store.Get(tt.id)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	store := NewStore()
	_ = store.Add(User{ID: "1", Name: "Alice"})

	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{"existing user", "1", nil},
		{"already deleted", "1", ErrUserNotFound},
		{"missing user", "999", ErrUserNotFound},
		{"empty id", "", ErrInvalidID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Delete(tt.id)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestList(t *testing.T) {
	store := NewStore()

	_ = store.Add(User{ID: "1", Name: "Alice"})
	_ = store.Add(User{ID: "2", Name: "Bob"})

	users := store.List()

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}
