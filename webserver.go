package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startServer(port string) {
	// Initialize the Gin router
	log.Println("Initializing Gin router...")
	router := gin.Default()

	// Load HTML templates
	router.Static("/api/v1/cssThemes", "./ThemeweaveBackend/cssThemes")   // Serve static CSS files
	router.Static("/api/v1/static", "./ThemeweaveBackend/library/static") // Serve static files
	router.LoadHTMLGlob("ThemeweaveBackend/templates/*")

	// Main entry point of website
	router.GET("/", HandleIndex)

	// --- API Routes ---
	// Group API routes under /api/v1
	v1 := router.Group("/api/v1")

	// Basic health check endpoint
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "ThemeWeave backend is running!",
		})
	})

	// if no port is specified, default to 8080
	log.Printf("ThemeWeave backend starting on :%s...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
