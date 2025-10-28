package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Title ThemeWeave
// Version 0.1
// Description: ThemeWeave builds static websites.
// Email: UndeadTokenArt@gmail.com
// License: MIT
// license: https://opensource.org/licenses/MIT
// BasePath /api/v1

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Ensure config file exists
	CheckConfig()

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// Start the web server
	// startServer(port)
}
