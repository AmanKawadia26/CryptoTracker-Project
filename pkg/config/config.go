package config

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Config struct {
	APIKey string `json:"api_key"`
}

var AppConfig Config

func LoadConfig() {
	// Attempt to open the configuration file
	file, err := os.Open("config.json")
	if err != nil {
		color.New(color.FgRed).Printf("Configuration file not found: %v\n", err)
		log.Fatalf("Configuration file not found: %v", err)
	}
	defer file.Close()

	// Attempt to decode the JSON configuration file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		color.New(color.FgRed).Printf("Error decoding configuration file: %v\n", err)
		log.Fatalf("Error decoding configuration file: %v", err)
	}

	// Success message for loading configuration
	color.New(color.FgGreen).Println("Configuration loaded successfully.")
}
