package main

import (
	"log"
	"net/http"

	"go-crud-api/config"
	"go-crud-api/routes"
)

func main() {
	// Initialize database connection
	config.ConnectDB()

	// Set up the router
	router := routes.SetupRouter()

	// Start the server
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
