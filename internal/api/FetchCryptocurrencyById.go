package api

import (
	"cryptotracker/models"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

// FetchCryptocurrencyById fetches a specific cryptocurrency by its ID from the CoinLore API
func FetchCryptocurrencyById(id string) (*models.Cryptocurrency, error) {
	url := fmt.Sprintf("https://api.coinlore.net/api/ticker/?id=%s", id)
	color.New(color.FgCyan).Printf("Fetching cryptocurrency data for ID: %s...\n", id)
	resp, err := http.Get(url)
	if err != nil {
		color.New(color.FgRed).Println("Error fetching data:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.New(color.FgRed).Println("Error reading response body:", err)
		return nil, err
	}

	var cryptoData []models.Cryptocurrency
	if err := json.Unmarshal(body, &cryptoData); err != nil {
		color.New(color.FgRed).Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	if len(cryptoData) > 0 {
		color.New(color.FgGreen).Printf("Data for cryptocurrency ID %s fetched successfully.\n", id)
		return &cryptoData[0], nil
	}
	color.New(color.FgYellow).Printf("Cryptocurrency with ID %s not found.\n", id)
	return nil, fmt.Errorf("cryptocurrency with id %s not found", id)
}
