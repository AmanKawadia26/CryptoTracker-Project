package api

import (
	"cryptotracker/models"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

// FetchCryptocurrencyData fetches the top cryptocurrencies data from the CoinLore API
func FetchCryptocurrencyData() ([]models.Cryptocurrency, error) {
	color.New(color.FgCyan).Println("Fetching top cryptocurrencies data...")
	resp, err := http.Get("https://api.coinlore.net/api/tickers/")
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

	var cryptoData struct {
		Data []models.Cryptocurrency `json:"data"`
	}
	if err := json.Unmarshal(body, &cryptoData); err != nil {
		color.New(color.FgRed).Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	color.New(color.FgGreen).Println("Data fetched successfully.")
	return cryptoData.Data, nil
}
