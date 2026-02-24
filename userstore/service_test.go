package userstore

import (
	"errors"
	"testing"
)

func TestService_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{"valid user", User{ID: "1", Name: "Alice", Age: 25}, nil},
		{"empty id", User{Name: "Bob", Age: 20}, ErrInvalidID},
		{"empty name", User{ID: "2", Age: 20}, ErrEmptyName},
		{"negative age", User{ID: "3", Name: "John", Age: -1}, ErrInvalidAge},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			repo := NewStore()
			service := NewService(repo)

			err := service.CreateUser(tt.user)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}
