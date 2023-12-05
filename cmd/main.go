package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"mvc-golang/internal/presentation/controller"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/users", controller.GetAllUsers)
	r.Get("/announcements", controller.GetAllAnnouncements)
	err := http.ListenAndServe(":5005", r)
	if err != nil {
		log.Fatal(err)
	}
}
