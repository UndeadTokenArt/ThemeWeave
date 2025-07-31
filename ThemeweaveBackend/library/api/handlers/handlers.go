package handlers

import (
	"fmt"
	"net/http"

	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/library/database"
	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	// Render the index page with a welcome message
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "Welcome to ThemeWeave",
		"Message": "This is the ThemeWeave backend. Use the API to manage your themes and websites.",
	})
}

// I should be using the context to pass the client ID, but for simplicity, I'm using a hardcoded value here.
func HandleClient(c *gin.Context) {

	// get the client ID from the context or request parameters
	dbID, err := c.Get("clientID")
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"Title":   "Bad Request",
			"Message": "Client ID is missing or invalid.",
		})
		return
	}

	client, err := database.GetWebsitefromDB(dbID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"Title":   "Internal Server Error",
			"Message": "Failed to retrieve client information.",
		})
		return
	}
	if client == nil {
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"Title":   "Client Not Found",
			"Message": "The requested client does not exist.",
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   client.Name,
		"Message": client.MainBody,
	})
}

// HandleCreateClient handles the creation of a new client (website) based on the provided json data.
func HandleCreateClient(c *gin.Context) {
	// Bind the incoming JSON to a Website struct
	var website database.Website

	// Validate the input
	if err := c.ShouldBindJSON(&website); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	// Save the new website to the database
	if err := database.DB.Create(&website).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create website: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Website created successfully", "website_id": website.ID})
}
