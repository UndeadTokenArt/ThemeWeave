package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/library/database"
	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	firstE := []string{"element1", "element2", "element3"}
	secE := []string{"itemA", "itemB", "itemC"}
	olList := [][]string{firstE, secE}

	c.HTML(http.StatusOK, "WebInterface.tmpl", gin.H{
		"Title":        database.Website{}.Name,
		"Message":      database.Website{}.HeaderContent,
		"Heirchierchy": olList,
		"Style":        template.HTML(database.Website{}.ColorScheme),
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
