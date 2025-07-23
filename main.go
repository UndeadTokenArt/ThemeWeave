package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title ThemeWeave API
// @version 0.0
// @description This is the backend API for the ThemeWeave website builder.
// @contact.email UndeadTokenArt@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Initialize the Gin router
	// Use gin.ReleaseMode() for production to disable debug output
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("ThemeweaveBackend/templates/*")

	// --- Middleware ---
	// Add any global middleware here, e.g., CORS, logging, authentication
	// For local development, you might need CORS to allow frontend to connect
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for development
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// --- API Routes ---
	// Group routes under a common prefix, e.g., /api/v1
	v1 := router.Group("/api/v1")
	{

		v1.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "WebInterface.tmpl", gin.H{
				"Title": "Welcome to the ThemeWeave API!",
			})
		})

		// Basic health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "ThemeWeave backend is running!",
			})
		})

		// Placeholder for future theme-related endpoints
		v1.GET("/themes", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Themes endpoint - coming soon!",
				"themes":  []string{"default", "minimal", "dark"}, // Example themes
			})
		})

		// Placeholder for future element-related endpoints
		v1.GET("/elements", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":  "Elements endpoint - coming soon!",
				"elements": []string{"text_block", "image_gallery", "button", "hero_section"}, // Example elements
			})
		})

		// You can add more routes here for:
		// - User authentication (login, register)
		// - Website project management (create, save, load)
		// - Theme customization options
		// - Element configuration
		// - Asset uploads
	}

	// --- Start the server ---
	// Listen and serve on port 8080
	log.Println("ThemeWeave backend starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
