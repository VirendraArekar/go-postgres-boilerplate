package routes

import (
	"crud/controllers"
	"crud/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", controllers.Register).Methods("POST")
	auth.HandleFunc("/login", controllers.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTAuth)
	api.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	api.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	api.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	api.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
