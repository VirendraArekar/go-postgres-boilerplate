package models

// User struct represents a user entity in the database
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TableName returns the table name for User model
func (u User) TableName() string {
	return "users"
}
