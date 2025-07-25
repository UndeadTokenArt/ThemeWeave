package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Database initialization and connection logic.
func InitDB() {
	db, err := gorm.Open(sqlite.Open("Themeweave.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	// AutoMigrate will create the table if it doesn't exist and update it if the schema changes
	err = DB.AutoMigrate(&Website{})
	if err != nil {
		log.Fatal(err)
	}
}

// GetDB returns the website data from the Database if the website exists
func GetWebsitefromDB(ClientID uint) (*Website, error) {
	var website Website

	result := DB.First(&website, ClientID)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get website: %v", result.Error)
	}
	return &website, nil
}
