package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/library/database"
)

func RunTests() {
	fmt.Println("Running tests...")

	// Test creating a website entry from JSON
	if err := CreateWebsiteEntry(); err != nil {
		fmt.Printf("Error creating website entry: %v\n", err)
	} else {
		fmt.Println("Website entry created successfully.")
	}
}

// CheckCreateWebsiteEntry checks if a website entry can be created from a JSON file
func CreateWebsiteEntry() error {
	// Read the JSON file
	jsonFile, err := os.Open("ThemeweaveBackend/tests/jsonTestfiles/DBClientinfo.json")
	if err != nil {
		return fmt.Errorf("error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Read the file content
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	// Create a Website struct to unmarshal into
	var website database.Website

	// Unmarshal JSON directly into the Website struct
	if err := json.Unmarshal(byteValue, &website); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	fmt.Printf("Successfully parsed website data: %+v\n", website)
	return nil
}
