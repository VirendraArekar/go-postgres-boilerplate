package controllers

import (
	"go-crud-api/services"
	"go-crud-api/utils"
	"net/http"
)

// GetUsers handles GET requests for all users
func GetAuth(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}
