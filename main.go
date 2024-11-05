package main

import (
	"crud/config"
	"crud/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
