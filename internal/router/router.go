package router

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	auth.RegisterRoutes(r)
	todo.RegisterRoutes(r)
	
	return r
}
