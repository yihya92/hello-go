package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *UserService
}

func NewHandler(service *UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAll()
	respondJSON(w, http.StatusOK, users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	user, err := h.service.GetByID(id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	if user.Name == "" || user.Email == "" {
		respondError(w, http.StatusBadRequest, "Name and Email required")
		return
	}

	created := h.service.Create(user)
	respondJSON(w, http.StatusCreated, created)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updated User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	user, err := h.service.Update(id, updated)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := h.service.Delete(id); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
