package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.jcosta86.com/todoapi/configs"
	"github.jcosta86.com/todoapi/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.ListAll)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)

	port := fmt.Sprintf(":%s", configs.GetServerPort())
	log.Printf("Server listening on port %s", port)
	http.ListenAndServe(port, r)
}
