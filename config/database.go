package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "user=postgres dbname=golang password=viren45mca sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		os.Exit(1)
	}

	// Ping the database to test connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
		os.Exit(1)
	}

	fmt.Println("Database connection established")
}
