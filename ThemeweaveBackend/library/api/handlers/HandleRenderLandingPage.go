package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Move struct definitions outside the function
type Client struct {
	CustomerID     string      `json:"customer_id"`
	ClientPortrait string      `json:"client_portrait"`
	HeroImage      string      `json:"hero_image"`
	Name           string      `json:"name"`
	Website        string      `json:"website"`
	ContactInfo    string      `json:"contact_info"`
	Status         string      `json:"status"`
	PaymentMethod  interface{} `json:"payment_method"`
	PaymentAmount  interface{} `json:"payment_amount"`
	TypeOfBusiness string      `json:"type_of_business"`
	Location       string      `json:"location"`
}

type ConfigFile struct {
	Clients []Client `json:"clients"`
}

func LoadConfigFile(c *gin.Context) ConfigFile {
	// Get current working directory properly
	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get working directory"})
		return ConfigFile{} // Return empty ConfigFile instead of nothing
	}

	// get data from config.json
	file, err := os.Open(wd + "/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open config file"})
		log.Printf("Error opening config file: %v at %s\n", err, wd+"/config.json")
		return ConfigFile{} // Return empty ConfigFile instead of nothing
	}
	defer file.Close()

	var configFile ConfigFile
	if err := json.NewDecoder(file).Decode(&configFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse config file"})
		log.Printf("Error decoding config file: %v at %s\n", err, wd+"/config.json")
		return ConfigFile{} // Return empty ConfigFile instead of nothing
	}
	return configFile
}

func prepareAssets(customerData *Client, wd string) {
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
	configFile := LoadConfigFile(c)
	if len(configFile.Clients) == 0 {
		return // LoadConfigFile already sent the error response
	}

	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get working directory"})
		return
	}
	log.Printf("Current working directory: %s\n", wd)

	CustomerID := c.Param("customer_id")

	// Find the customer by ID
	var customerData *Client
	for _, client := range configFile.Clients {
		if client.CustomerID == CustomerID {
			customerData = &client
			break
		}
	}

	// Check if customer data was found
	if customerData == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	fmt.Printf("Rendering landing page with data: %+v\n", *customerData)

	// Create a buffer to capture the rendered HTML
	tmpl, err := template.ParseFiles("ThemeweaveBackend/templates/landingpage.tmpl")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template"})
		log.Printf("Error parsing template: %v\n", err)
		return
	}

	// Render the template with the customer data
	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, *customerData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		log.Printf("Error executing template: %v\n", err)
		log.Printf("Customer data being passed: %+v\n", *customerData)
		return
	}

	// export static page to public/customerName
	// create dir if customerName dir does not exist
	customerDir := fmt.Sprintf("./public/%s", customerData.Name)
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
	prepareAssets(customerData, wd)

	// Serve the static rendered page
	c.File(htmlFile)
}
