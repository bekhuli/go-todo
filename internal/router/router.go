package router

import (
	"github.com/bekhuli/go-todo/internal/auth"
	"github.com/bekhuli/go-todo/internal/todo"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	auth.RegisterRoutes(r)
	todo.RegisterRoutes(r)

	return r
}
