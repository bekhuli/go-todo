package main

import (
	"log"
	"net/http"

	"github.com/bekhuli/go-todo/config"
	"github.com/bekhuli/go-todo/internal/db"
	"github.com/bekhuli/go-todo/internal/router"
)

func main() {
	cfg := config.Envs

	db.Init(cfg.DatabaseURL)

	r := router.NewRouter()

	log.Println("Service is running on port:", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
