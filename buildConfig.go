package main

import (
	"log"
	"os"
)

func CheckConfig() {
	data := []byte(`  {
    "clients": [
      {
        "customer_id": "1",
        "client_portrait": "images/logo.png",
        "hero_image": "images/hero.jpg",
        "name": "Sample Client",
        "website": "images/logo.png",
        "contact_info": "15551234567",
        "template": "default",
        "location": "1234 N Main St. City, ST 12345",
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
