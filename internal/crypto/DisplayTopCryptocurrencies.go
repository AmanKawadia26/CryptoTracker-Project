package crypto

import (
	"cryptotracker/internal/api"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
)

func DisplayTopCryptocurrencies() {
	params := map[string]string{
		"start":   "1",
		"limit":   "10",
		"convert": "USD",
	}

	response := api.GetAPIResponse("/listings/latest", params)

	var result map[string]interface{}
	json.Unmarshal(response, &result)

	data := result["data"].([]interface{})

	color.New(color.FgGreen).Println("Top 10 Cryptocurrencies:")
	for i, crypto := range data {
		cryptoMap := crypto.(map[string]interface{})
		name := cryptoMap["name"].(string)
		symbol := cryptoMap["symbol"].(string)
		price := cryptoMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["price"].(float64)

		fmt.Printf("%d. %s (%s): $%.2f\n", i+1, name, symbol, price)
	}
	fmt.Println()
}
