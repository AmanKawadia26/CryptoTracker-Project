package crypto

import (
	"cryptotracker/internal/api"
	"cryptotracker/models"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"strings"
	"time"
)

func SearchCryptocurrency() {
	var input string
	color.New(color.FgCyan).Print("Enter the symbol or name of the cryptocurrency: ")
	fmt.Scan(&input)

	// Normalize the input to lowercase for case-insensitive comparison
	input = strings.ToLower(input)

	// Parameters for the API request to fetch all cryptocurrencies
	params := map[string]string{
		"start":   "1",
		"limit":   "5000", // Load up to 5000 cryptocurrencies for now
		"convert": "USD",
	}

	// Make a single API call to get all the cryptocurrencies
	response := api.GetAPIResponse("/listings/latest", params)

	var result map[string]interface{}
	err := json.Unmarshal(response, &result)
	if err != nil {
		color.New(color.FgRed).Printf("Error unmarshalling API response: %v\n", err)
		return
	}

	// Extract data from the API response
	data, ok := result["data"].([]interface{})
	if !ok {
		color.New(color.FgRed).Println("Data not found in the response.")
		return
	}

	// Iterate over the list of cryptocurrencies and look for a match by symbol or name
	for _, item := range data {
		crypto := item.(map[string]interface{})

		// Get the symbol and name, normalize to lowercase for comparison
		symbol := strings.ToLower(crypto["symbol"].(string))
		name := strings.ToLower(crypto["name"].(string))

		// Check if input matches either the symbol or name
		if symbol == input || name == input {
			// Proceed with normal processing if the cryptocurrency is found
			priceObj, ok := crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["price"]
			if !ok {
				color.New(color.FgRed).Println("Could not retrieve the price of the cryptocurrency.")
				return
			}

			price, ok := priceObj.(float64)
			if !ok {
				color.New(color.FgRed).Println("Failed to convert price to float64.")
				return
			}

			// Display the result
			color.New(color.FgGreen).Printf("%s (%s): $%.2f\n", crypto["name"].(string), crypto["symbol"].(string), price)
			fmt.Println()

			// Generate and display the graph
			displayCryptoGraph(crypto["name"].(string), price)
			return
		}
	}

	// If no match is found
	color.New(color.FgYellow).Printf("Cryptocurrency not found for input: %s\n", input)
	color.New(color.FgMagenta).Println("Please request the addition of this cryptocurrency to our app.")

	// Create a new request for the unavailable cryptocurrency
	request := &models.UnavailableCryptoRequest{
		CryptoSymbol:   input,
		RequestMessage: "Please add this cryptocurrency.",
		Status:         "Pending",
	}

	// Save the request to the file
	if err := saveUnavailableCryptoRequest(request); err != nil {
		color.New(color.FgRed).Printf("Error saving unavailable crypto request: %v\n", err)
		return
	}
	color.New(color.FgGreen).Println("Request to add the cryptocurrency has been submitted.")
}

func displayCryptoGraph(cryptoName string, currentPrice float64) {
	// Generate random price data for the last 30 days
	prices := generateRandomPrices(30, currentPrice)

	// Display the graph
	color.New(color.FgCyan).Printf("30-day price graph for %s:\n\n", cryptoName)

	maxPrice := prices[0]
	minPrice := prices[0]
	for _, price := range prices {
		if price > maxPrice {
			maxPrice = price
		}
		if price < minPrice {
			minPrice = price
		}
	}

	graphHeight := 20
	for i := 0; i < graphHeight; i++ {
		price := maxPrice - (float64(i) * (maxPrice - minPrice) / float64(graphHeight-1))
		fmt.Printf("%8.2f |", price)

		for _, p := range prices {
			if p >= price {
				color.New(color.FgGreen).Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Print("         ")
	fmt.Println(strings.Repeat("-", len(prices)))
	fmt.Print("         ")
	for i := 0; i < len(prices); i += len(prices) / 5 {
		fmt.Printf("%-6d", 30-i)
	}
	fmt.Println("\n         Days ago")
}

func generateRandomPrices(days int, currentPrice float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	prices := make([]float64, days)
	prices[days-1] = currentPrice

	for i := days - 2; i >= 0; i-- {
		// Generate a random percentage change between -5% and 5%
		change := (rand.Float64() - 0.5) * 0.1
		prices[i] = prices[i+1] * (1 + change)
	}

	return prices
}
