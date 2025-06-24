
package main

import (
	"log"

	"hexagonal-backend/internal/adapter/rest"
	"hexagonal-backend/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize and start the REST server
	server := rest.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
