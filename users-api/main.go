package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	service := NewUserService()    // handles business logic
	handler := NewHandler(service) // handles at the http layer
	r := mux.NewRouter()           //creates a new HTTP request router

	r.HandleFunc("/users", handler.GetUsers).Methods("GET")           // get method for all users
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")       // get method for specific user
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")        // put method to create new user
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")    // put method to update specific user
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE") // delete method to delete specific user

	// Register middlewares
	r.Use(recoveryMiddleware)
	r.Use(requestIDMiddleware)
	r.Use(loggingMiddleware)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":9001", r))

}

// var (
// 	users  = make(map[int]User)
// 	nextID = 1
// 	mu     sync.Mutex
// )

// // GET /users
// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	list := []User{} // empty slice
// 	for _, user := range users {
// 		list = append(list, user)
// 	}
// 	respondJSON(w, http.StatusOK, list)
// }

// // GET /users/{id}
// func getUser(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	params := mux.Vars(r)               // extracts URL path parameters from the request
// 	id, _ := strconv.Atoi(params["id"]) // Atoi means: ASCII to Integer

// 	user, exists := users[id]
// 	if !exists {
// 		respondError(w, http.StatusNotFound, "User Not Found")
// 		return
// 	}

// 	respondJSON(w, http.StatusOK, user)
// }

// // POST /users
// func createUser(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		respondError(w, http.StatusNotFound, "Invalid Input")
// 		return
// 	}
// 	if user.Name == "" || user.Email == "" {
// 		respondError(w, http.StatusNotFound, "Name and Email Required")
// 		return
// 	}
// 	user.ID = nextID
// 	nextID++
// 	users[user.ID] = user

// 	respondJSON(w, http.StatusCreated, user)

// }

// // PUT /users/{id}
// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	params := mux.Vars(r)
// 	id, _ := strconv.Atoi(params["id"])
// 	_, exists := users[id]
// 	if !exists {
// 		respondError(w, http.StatusNotFound, "User not found")
// 		return
// 	}
// 	var updated User
// 	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
// 		respondError(w, http.StatusNotFound, "Invalid Input")
// 		return
// 	}
// 	updated.ID = id
// 	users[id] = updated

// 	respondJSON(w, http.StatusOK, updated)
// }

// // DELETE /users/{id}
// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	params := mux.Vars(r)
// 	id, _ := strconv.Atoi(params["id"])

// 	_, exists := users[id]
// 	if !exists {
// 		respondError(w, http.StatusNotFound, "User not found")
// 		return
// 	}
// 	delete(users, id)
// 	w.WriteHeader(http.StatusNoContent)
// }

// respond json method used to save code
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := ErrorResponse{
		Error: ErrorDetail{
			Code:    status,
			Message: message,
		},
	}

	json.NewEncoder(w).Encode(response)
}

/*
main.go	App bootstrap + routes
models.go	Structs
service.go	Business logic
handlers.go	HTTP layer
middleware.go	Cross-cutting concerns
*/
