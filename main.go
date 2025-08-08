package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/library/api/handlers"
	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/library/database"
	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/tests"
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
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Initialize the database connection
	log.Println("Initializing database connection...")
	database.InitDB()

	// Initialize the Gin router
	log.Println("Initializing Gin router...")
	router := gin.Default()

	// Load HTML templates
	router.Static("/api/v1/cssThemes", "./ThemeweaveBackend/cssThemes")   // Serve static CSS files
	router.Static("/api/v1/static", "./ThemeweaveBackend/library/static") // Serve static files
	router.LoadHTMLGlob("ThemeweaveBackend/templates/*")

	// --- Middleware ---
	// Add any global middleware here, e.g., CORS, logging, authentication
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

	// --- check for testing flag  ---
	runTests := flag.Bool("testing", false, "Run tests before starting the server")
	flag.Parse()

	if *runTests {
		log.Println("Running tests...")
		tests.RunTests()
	} else {
		log.Println("Skipping tests, starting server...")
	}

	// --- testing endpoints ---
	testing := router.Group("/testing")                        // Testing group for development purposes
	testing.POST("/createClient", handlers.HandleCreateClient) // Testing endpoint for creating a new client (website)

	// This is the main page of the website, it will be served at the root path
	router.GET("/", handlers.HandleIndex)
	router.GET("/renderCustomer/:customer_id", handlers.HandleLandingPage) // Redirect /index to the main page

	// Contact form submission
	router.POST("/contact", handlers.HandleContactForm)

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

	// --- Start the server ---
	// Listen and serve on port 8080
	log.Println("ThemeWeave backend starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
