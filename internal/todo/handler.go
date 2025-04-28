package todo

import (
	"encoding/json"
	"github.com/bekhuli/go-todo/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TodoRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/todos").Subrouter()
	s.Use(middleware.JWTMiddleware)
	s.HandleFunc("", CreateHandler).Methods("POST")
	s.HandleFunc("", ListHandler).Methods("GET")
	s.HandleFunc("/{id}", UpdateHandler).Methods("PUT")
	s.HandleFunc("/{id}", DeleteHandler).Methods("DELETE")
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserKey).(int)

	var req TodoRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := Create(userID, req.Title)
	if err != nil {
		http.Error(w, "Failed to create todos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserKey).(int)

	todos, err := List(userID)
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(todos)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var req TodoRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := Update(id, req.Title, req.Done)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := DeleteTodo(id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
