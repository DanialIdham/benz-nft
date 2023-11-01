package database

import (
	"fmt"

	"github.com/danial.idham/benz/pkg/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New initializes and returns a new gorm.DB instance connected to the provided DSN.
// It will also run migrations for the Profile model.
// Any error in the process will be returned.
func New(dsn string) (*gorm.DB, error) {
	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Migrate the schema for the Profile model
	if err := db.AutoMigrate(&repo.Profile{}); err != nil {
		return nil, fmt.Errorf("failed to migrate Profile schema: %w", err)
	}

	return db, nil
}
