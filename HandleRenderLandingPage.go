package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func prepareAssets(customerData *ConfigData, wd string) {
	staticDir := fmt.Sprintf("%s/ThemeweaveBackend/library/static", wd)

	// Move the images and other assets to the public/customerName Directory
	customerDir := fmt.Sprintf("%s/public/%s", wd, customerData.Name)
	if err := os.MkdirAll(customerDir, os.ModePerm); err != nil {
		log.Printf("Error creating directory: %v\n", err)
		return
	}

	// Check if hero image exists before moving
	if customerData.HeroImage != "" {
		heroImageSrc := fmt.Sprintf("%s/%s", staticDir, customerData.HeroImage)
		heroImageDest := fmt.Sprintf("%s/%s", customerDir, customerData.HeroImage)

		// Check if source file exists
		if _, err := os.Stat(heroImageSrc); os.IsNotExist(err) {
			log.Printf("Hero image source file does not exist: %s\n", heroImageSrc)
		} else {
			// Create destination directory if it doesn't exist (for subdirectories like images/)
			destDir := filepath.Dir(heroImageDest)
			if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
				log.Printf("Error creating destination directory: %v\n", err)
				return
			}

			// Copy instead of move to avoid losing the original
			if err := copyFile(heroImageSrc, heroImageDest); err != nil {
				log.Printf("Error copying hero image: %v\n", err)
			}
		}
	}

	// Check if client portrait exists before moving
	if customerData.ClientPortrait != "" {
		clientPortraitSrc := fmt.Sprintf("%s/%s", staticDir, customerData.ClientPortrait)
		clientPortraitDest := fmt.Sprintf("%s/%s", customerDir, customerData.ClientPortrait)

		// Check if source file exists
		if _, err := os.Stat(clientPortraitSrc); os.IsNotExist(err) {
			log.Printf("Client portrait source file does not exist: %s\n", clientPortraitSrc)
		} else {
			// Create destination directory if it doesn't exist
			destDir := filepath.Dir(clientPortraitDest)
			if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
				log.Printf("Error creating destination directory: %v\n", err)
				return
			}

			// Copy instead of move to avoid losing the original
			if err := copyFile(clientPortraitSrc, clientPortraitDest); err != nil {
				log.Printf("Error copying client portrait: %v\n", err)
			}
		}
	}
}

// Helper function to copy files
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// HandleLandingPage renders the landing page for a specific customer based on their ID.
func HandleLandingPage(c *gin.Context) {
	// Load config file
	siteData := LoadConfigFile(c)

	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get working directory"})
		return
	}
	log.Printf("Current working directory: %s\n", wd)
	log.Panicln("Rendering landing page with data")

	// Create a buffer to capture the rendered HTML
	tmpl, err := template.ParseFiles(siteData.Template + ".tmpl")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template"})
		log.Printf("Error parsing template: %v\n", err)
		return
	}

	// Render the template with the customer data
	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, siteData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		log.Printf("Error executing template: %v\n", err)
		log.Printf("Customer data being passed: %+v\n", siteData)
		return
	}

	// export static page to public/customerName
	// create dir if customerName dir does not exist
	customerDir := fmt.Sprintf("./public/%s", siteData.Name)
	if err := os.MkdirAll(customerDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create directory: %v", err)})
		return
	}

	// Save the rendered HTML to a file
	htmlFile := fmt.Sprintf("%s/landingpage.html", customerDir)
	if err := os.WriteFile(htmlFile, htmlBuffer.Bytes(), os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save HTML file: %v", err)})
		return
	}

	// move the images and other assets to the public/customerName Directory
	prepareAssets(&siteData, wd)

	// Serve the static rendered page
	c.File(htmlFile)
}
