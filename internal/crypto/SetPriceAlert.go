package crypto

import (
	"cryptotracker/internal/api"
	"cryptotracker/models"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	//"io/ioutil"
	"time"
)

func SetPriceAlert(user *models.User) {
	var symbol string
	var targetPrice float64

	color.New(color.FgCyan).Print("Enter the symbol of the cryptocurrency: ")
	fmt.Scan(&symbol)
	color.New(color.FgCyan).Print("Enter your target price in USD: ")
	fmt.Scan(&targetPrice)

	params := map[string]string{
		"symbol":  symbol,
		"convert": "USD",
	}

	response := api.GetAPIResponse("/quotes/latest", params)

	var result map[string]interface{}
	err := json.Unmarshal(response, &result)
	if err != nil {
		color.New(color.FgRed).Printf("Error unmarshalling API response: %v\n", err)
		return
	}

	// Check if the data for the symbol is available
	data, dataOk := result["data"].(map[string]interface{})
	if !dataOk || data[symbol] == nil {
		color.New(color.FgRed).Printf("Cryptocurrency data not found for symbol: %s\n", symbol)
		return
	}

	// Proceed if the cryptocurrency data is found
	cryptoData, ok := data[symbol].(map[string]interface{})
	if !ok {
		color.New(color.FgRed).Printf("Unexpected data structure for symbol: %s\n", symbol)
		return
	}

	// Safely check for price data in the response
	quote, ok := cryptoData["quote"].(map[string]interface{})
	if !ok || quote["USD"] == nil {
		color.New(color.FgRed).Printf("Quote data not available for %s\n", symbol)
		return
	}

	priceData, ok := quote["USD"].(map[string]interface{})
	if !ok || priceData["price"] == nil {
		color.New(color.FgRed).Printf("Price data not available for %s\n", symbol)
		return
	}

	cryptoIDInterface, ok := cryptoData["id"].(float64) // Adjust based on actual data type of ID
	if !ok || cryptoIDInterface == 0 {
		color.New(color.FgRed).Printf("Cryptocurrency ID not found for symbol: %s\n", symbol)
		return
	}

	cryptoID := int(cryptoIDInterface)

	currentPrice, ok := priceData["price"].(float64)
	if !ok {
		color.New(color.FgRed).Println("Failed to convert price to float64.")
		return
	}

	// Check if the target price has been met
	if currentPrice >= targetPrice {
		color.New(color.FgGreen).Printf("Alert: %s has reached your target price of $%.2f. Current price: $%.2f\n", symbol, targetPrice, currentPrice)
	} else {
		// Create a notification request if the target price is not met
		notification := &models.PriceNotification{
			CryptoID:    cryptoID,
			Crypto:      symbol,
			TargetPrice: targetPrice,
			Username:    user.Username,
			AskedAt:     time.Now().Format(time.RFC3339),
			Status:      "Pending",
		}

		// Save the notification
		err = savePriceNotification(notification)
		if err != nil {
			color.New(color.FgRed).Printf("Error saving notification: %v\n", err)
			return
		}

		color.New(color.FgYellow).Printf("%s is still below your target price. Current price: $%.2f. Notification created.\n", symbol, currentPrice)
	}
}
