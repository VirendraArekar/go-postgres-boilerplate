package services

import (
	"go-crud-api/config"
	"go-crud-api/models"
	"log"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Printf("Error scanning user row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user models.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

// GetUser retrieves a user by ID
func GetUser(id int) (models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return user, err
	}
	return user, nil
}

// UpdateUser updates user details in the database
func UpdateUser(user models.User) error {
	_, err := config.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// DeleteUser removes a user from the database
func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
