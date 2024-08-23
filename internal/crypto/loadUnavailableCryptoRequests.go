package crypto

import (
	"cryptotracker/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadUnavailableCryptoRequests() ([]*models.UnavailableCryptoRequest, error) {
	if _, err := os.Stat("unavailable_cryptos.json"); os.IsNotExist(err) {
		return []*models.UnavailableCryptoRequest{}, nil
	}

	data, err := ioutil.ReadFile("unavailable_cryptos.json")
	if err != nil {
		return nil, err
	}

	var requests []*models.UnavailableCryptoRequest
	if err := json.Unmarshal(data, &requests); err != nil {
		return nil, err
	}

	return requests, nil
}
