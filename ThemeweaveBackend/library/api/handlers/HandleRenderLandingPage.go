package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// HandleLandingPage renders the landing page for a specific customer based on their ID.
func HandleLandingPage(c *gin.Context) {
	// Get current working directory properly
	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get working directory"})
		return
	}

	log.Printf("Current working directory: %s\n", wd)

	CustomerID := c.Param("customer_id")

	// get data from config.json
	file, err := os.Open(wd + "/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open config file"})
		log.Printf("Error opening config file: %v at %s\n", err, wd+"/config.json")
		return
	}
	defer file.Close()

	// Update struct to match your JSON structure
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

	var configFile ConfigFile
	if err := json.NewDecoder(file).Decode(&configFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse config file"})
		log.Printf("Error decoding config file: %v at %s\n", err, wd+"/config.json")
		return
	}

	// Find the customer by ID
	var customerData *Client
	for _, client := range configFile.Clients {
		if client.CustomerID == CustomerID {
			customerData = &client
			break
		}
	}

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

	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, *customerData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		log.Printf("Error executing template: %v\n", err)
		log.Printf("Customer data being passed: %+v\n", *customerData)
		return
	}

	// save rendered html to public directory/customerID
	// create dir if customerId dir does not exist
	customerDir := fmt.Sprintf("./public/%s", CustomerID)
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

	// Send the response to the client
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlBuffer.Bytes())
}
