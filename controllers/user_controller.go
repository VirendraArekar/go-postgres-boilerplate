package controllers

import (
	"encoding/json"
	"go-crud-api/models"
	"go-crud-api/services"
	"go-crud-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers handles GET requests for all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}

// GetUser handles GET requests for a specific user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	user, err := services.GetUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "User not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// CreateUser handles POST requests to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// UpdateUser handles PUT requests to update user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = id

	err := services.UpdateUser(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// DeleteUser handles DELETE requests to remove a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := services.DeleteUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted"})
}
