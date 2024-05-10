package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alanpramil7/go-auth-service/internal/api"
	"github.com/alanpramil7/go-auth-service/internal/config"
	"github.com/alanpramil7/go-auth-service/pkg/db"
)

func main() {
	//Load config
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	// Initialize Database
	db, err := db.Initialize(config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Error initiaizing db: %v", err)
	}

	defer db.Close()

	// Setup router
	router := api.NewRouter(db)
	serverAddress := config.ServerAddress
	fmt.Println("Server started on post 8080")
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("Error on setting up the  server: %v", err)
	}
}
