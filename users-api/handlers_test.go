package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func setupTestRouter() (*mux.Router, *UserService) {
	service := NewUserService()
	handler := NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")

	return r, service
}

func TestCreateUserHandler(t *testing.T) {
	router, _ := setupTestRouter()

	payload := map[string]string{
		"name":  "Alice",
		"email": "alice@example.com",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rr.Code)
	}
}

func TestGetUsersHandler(t *testing.T) {
	router, _ := setupTestRouter()

	req := httptest.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestGetUserHandler_Success(t *testing.T) {
	router, service := setupTestRouter()

	created := service.Create(User{
		Name:  "Bob",
		Email: "bob@example.com",
	})

	req := httptest.NewRequest("GET", "/users/"+strconv.Itoa(created.ID), nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestGetUserHandler_NotFound(t *testing.T) {
	router, _ := setupTestRouter()

	req := httptest.NewRequest("GET", "/users/999", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rr.Code)
	}
}

func TestUpdateUserHandler(t *testing.T) {
	router, service := setupTestRouter()

	created := service.Create(User{
		Name:  "Old",
		Email: "old@example.com",
	})

	payload := map[string]string{
		"name":  "New",
		"email": "new@example.com",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("PUT", "/users/"+strconv.Itoa(created.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestDeleteUserHandler(t *testing.T) {
	router, service := setupTestRouter()

	created := service.Create(User{
		Name:  "DeleteMe",
		Email: "delete@example.com",
	})

	req := httptest.NewRequest("DELETE", "/users/"+strconv.Itoa(created.ID), nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rr.Code)
	}
}
