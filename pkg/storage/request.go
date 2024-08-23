package storage

import (
	"cryptotracker/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

const requestFile = "requests.json"

// GetAllUsers retrieves all users
func GetAllUsers() ([]*models.User, error) {
	data, err := ioutil.ReadFile("users.json") 
	if err != nil {
		color.New(color.FgRed).Printf("Error reading user file: %v\n", err)
		log.Println("Error reading user file:", err)
		return nil, err
	}

	var users []*models.User
	if err := json.Unmarshal(data, &users); err != nil {
		color.New(color.FgRed).Printf("Error unmarshalling user data: %v\n", err)
		log.Println("Error unmarshalling user data:", err)
		return nil, err
	}

	color.New(color.FgGreen).Println("Users retrieved successfully.")
	return users, nil
}

// GetAllRequests retrieves all cryptocurrency requests
func GetAllRequests() ([]*models.Request, error) {
	data, err := ioutil.ReadFile(requestFile)
	if err != nil {
		color.New(color.FgRed).Printf("Error reading request file: %v\n", err)
		log.Println("Error reading request file:", err)
		return nil, err
	}

	var requests []*models.Request
	if err := json.Unmarshal(data, &requests); err != nil {
		color.New(color.FgRed).Printf("Error unmarshalling request data: %v\n", err)
		log.Println("Error unmarshalling request data:", err)
		return nil, err
	}

	color.New(color.FgGreen).Println("Requests retrieved successfully.")
	return requests, nil
}

// loadRequests loads cryptocurrency requests from the file
func loadRequests() ([]*models.Request, error) {
	if _, err := os.Stat(requestFile); os.IsNotExist(err) {
		color.New(color.FgYellow).Println("Request file not found, initializing a new one.")
		return []*models.Request{}, nil
	}

	data, err := ioutil.ReadFile(requestFile)
	if err != nil {
		color.New(color.FgRed).Printf("Error reading request file: %v\n", err)
		return nil, err
	}

	var requests []*models.Request
	if err := json.Unmarshal(data, &requests); err != nil {
		color.New(color.FgRed).Printf("Error unmarshalling request data: %v\n", err)
		return nil, err
	}

	return requests, nil
}

// UpdateRequestStatus updates the status of a specific cryptocurrency request
func UpdateRequestStatus(request *models.Request) error {
	requests, err := loadRequests()
	if err != nil {
		color.New(color.FgRed).Printf("Error loading requests: %v\n", err)
		return err
	}

	updated := false
	for i, r := range requests {
		if r.ID == request.ID { 
			requests[i] = request
			updated = true
			break
		}
	}

	if !updated {
		color.New(color.FgYellow).Printf("Request with ID %s not found.\n", request.ID)
		return nil
	}

	data, err := json.Marshal(requests)
	if err != nil {
		color.New(color.FgRed).Printf("Error marshaling request data: %v\n", err)
		return err
	}

	if err := ioutil.WriteFile(requestFile, data, 0644); err != nil {
		color.New(color.FgRed).Printf("Error writing request data to file: %v\n", err)
		return err
	}

	color.New(color.FgGreen).Println("Request status updated successfully.")
	return nil
}

// SaveRequest saves a cryptocurrency request
func SaveRequest(request *models.Request) error {
	requests, err := loadRequests()
	if err != nil {
		color.New(color.FgRed).Printf("Error loading requests: %v\n", err)
		return err
	}

	requests = append(requests, request)
	data, err := json.Marshal(requests)
	if err != nil {
		color.New(color.FgRed).Printf("Error marshaling request data: %v\n", err)
		return err
	}

	if err := ioutil.WriteFile(requestFile, data, 0644); err != nil {
		color.New(color.FgRed).Printf("Error writing request data to file: %v\n", err)
		return err
	}

	color.New(color.FgGreen).Println("Request saved successfully.")
	return nil
}
