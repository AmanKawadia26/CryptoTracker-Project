package storage

import (
	"cryptotracker/models"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

const unavailableCryptoRequestsFile = "unavailable_cryptos.json"

// SaveUnavailableCryptoRequest saves a new unavailable crypto request
func SaveUnavailableCryptoRequest(request *models.UnavailableCryptoRequest) error {
	requests, err := loadUnavailableCryptoRequests()
	if err != nil {
		color.Red("Failed to load unavailable crypto requests: %v", err)
		return err
	}

	requests = append(requests, request)
	data, err := json.Marshal(requests)
	if err != nil {
		color.Red("Failed to marshal unavailable crypto requests: %v", err)
		return err
	}

	if err := ioutil.WriteFile(unavailableCryptoRequestsFile, data, 0644); err != nil {
		color.Red("Failed to write unavailable crypto requests file: %v", err)
		return err
	}

	color.Green("Unavailable crypto request saved successfully!")
	return nil
}

// GetAllUnavailableCryptoRequests retrieves all unavailable crypto requests
func GetAllUnavailableCryptoRequests() ([]*models.UnavailableCryptoRequest, error) {
	data, err := ioutil.ReadFile(unavailableCryptoRequestsFile)
	if err != nil {
		color.Red("Error reading unavailable crypto requests file: %v", err)
		return nil, err
	}

	var requests []*models.UnavailableCryptoRequest
	if err := json.Unmarshal(data, &requests); err != nil {
		color.Red("Error unmarshalling unavailable crypto requests data: %v", err)
		return nil, err
	}

	color.Green("Retrieved all unavailable crypto requests successfully!")
	return requests, nil
}

// loadUnavailableCryptoRequests loads unavailable crypto requests from the file
func loadUnavailableCryptoRequests() ([]*models.UnavailableCryptoRequest, error) {
	if _, err := os.Stat(unavailableCryptoRequestsFile); os.IsNotExist(err) {
		color.Yellow("Unavailable crypto requests file does not exist. Creating a new one.")
		return []*models.UnavailableCryptoRequest{}, nil
	}

	data, err := ioutil.ReadFile(unavailableCryptoRequestsFile)
	if err != nil {
		color.Red("Failed to read unavailable crypto requests file: %v", err)
		return nil, err
	}

	var requests []*models.UnavailableCryptoRequest
	if err := json.Unmarshal(data, &requests); err != nil {
		color.Red("Failed to unmarshal unavailable crypto requests: %v", err)
		return nil, err
	}

	color.Green("Loaded unavailable crypto requests successfully!")
	return requests, nil
}
