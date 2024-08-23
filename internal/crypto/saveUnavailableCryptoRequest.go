package crypto

import (
	"cryptotracker/models"
	"encoding/json"
	"io/ioutil"
)

func saveUnavailableCryptoRequest(request *models.UnavailableCryptoRequest) error {
	requests, err := loadUnavailableCryptoRequests()
	if err != nil {
		return err
	}

	requests = append(requests, request)
	data, err := json.Marshal(requests)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("unavailable_cryptos.json", data, 0644); err != nil {
		return err
	}

	return nil
}
