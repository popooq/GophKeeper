package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gtihub.com/popooq/Gophkeeper/server/internal/config"
	"gtihub.com/popooq/Gophkeeper/server/internal/database"
	"gtihub.com/popooq/Gophkeeper/server/internal/handlers"
)

func main() {
	context := context.Background()
	config := config.New()
	database := database.New(context, config.DatabaseAddress)
	database.Migrate()

	handlers := handlers.New(database)
	router := chi.NewRouter()
	router.Mount("/", handlers.Route())

	server := http.Server{
		Addr:    config.Address,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
