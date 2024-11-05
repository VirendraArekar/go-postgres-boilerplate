package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	absPath, err := filepath.Abs("database/migrations")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	dbURL := "postgres://username:password@localhost:5432/database?sslmode=disable" // Replace with your database URL
	source := "file://" + absPath                                                   // Use absolute path for migrations

	// Parse command-line arguments
	action := flag.String("action", "up", "Specify migration action: up, down, drop")
	steps := flag.Int("steps", 1, "Specify number of steps for down action")
	flag.Parse()

	// Initialize migrate instance
	m, err := migrate.New(source, dbURL)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	// Handle migration actions
	switch *action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration up error: %v", err)
		}
		fmt.Println("Migration applied successfully.")

	case "down":
		if *steps > 0 {
			if err := m.Steps(-*steps); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Migration down error: %v", err)
			}
		} else {
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Migration down error: %v", err)
			}
		}
		fmt.Println("Migration rolled back successfully.")

	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatalf("Migration drop error: %v", err)
		}
		fmt.Println("All migrations dropped successfully.")

	default:
		fmt.Println("Invalid action. Use -action=up, -action=down, or -action=drop.")
		os.Exit(1)
	}

	// Close the migration connection
	srcErr, dbErr := m.Close()
	if srcErr != nil || dbErr != nil {
		log.Fatalf("Could not close migration resources: %v %v", srcErr, dbErr)
	}
}
