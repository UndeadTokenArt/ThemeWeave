package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckConfig() {
	data := []byte(`  {
    "clients": [
      {
        "client_portrait": "images/logo.png",
        "hero_image": "images/hero.jpg",
        "name": "Sample Client",
        "website": "images/logo.png",
        "contact_info": "15551234567",
        "template": "default",
        "location": "1234 N Main St. City, ST 12345"
      }
    ]
  }`)

	// Read the config.json file
	configFile, err := os.ReadFile("config.json")
	if !os.IsExist(err) || len(configFile) == 0 {
		log.Println("Creating default config file...")
		os.WriteFile("config.json", data, 0644)
	}
	if len(configFile) > 0 {
		log.Println("Config file found and is not empty.")
	}
}

func LoadConfigFile(c *gin.Context) ConfigData {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get working directory"})
		return ConfigData{} // Return empty ConfigFile instead of nothing
	}

	// get data from config.json
	file, err := os.Open(wd + "/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open config file"})
		log.Printf("Error opening config file: %v at %s\n", err, wd+"/config.json")
		return ConfigData{} // Return empty ConfigFile instead of nothing
	}
	defer file.Close()

	// Decode JSON data and assign to configFile
	var configFile ConfigData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode config file"})
		log.Printf("Error decoding config file: %v\n", err)
		return ConfigData{} // Return empty ConfigFile instead of nothing
	}
	return configFile
}
