package crypto

import (
	"cryptotracker/internal/api"
	"cryptotracker/models"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"strconv"
	"time"
)

// CheckNotifications checks if any notification requests for a specific user have met their target price
func CheckNotifications(username string) {
	notifications, err := LoadPriceNotifications()
	if err != nil {
		color.New(color.FgRed).Printf("Error loading notifications: %v\n", err)
		return
	}

	// Filter notifications by username
	userNotifications := filterNotificationsByUsername(notifications, username)

	// Iterate over all saved notifications for the user
	for _, notification := range userNotifications {
		if notification.Status == "Pending" {
			// Check current price for the cryptocurrency using its ID
			params := map[string]string{
				"id":     strconv.Itoa(notification.CryptoID), // Assuming API accepts ID as a string
				"convert": "USD",
			}

			// Assuming this is inside your CheckNotifications function
			response := api.GetAPIResponse("/info", params)

			var result map[string]interface{}
			err := json.Unmarshal(response, &result)
			if err != nil {
   				color.New(color.FgRed).Printf("Error unmarshalling response: %v\n", err)
    			return
			}

			// Check if the result is nil before proceeding
			if result == nil {
    			color.New(color.FgRed).Println("No data received from API.")
    			return
			}

			// Now safely proceed with type assertions knowing result is not nil
			data := result["data"].(map[string]interface{})[notification.Crypto].(map[string]interface{})
			price := data["quote"].(map[string]interface{})["USD"].(map[string]interface{})["price"].(float64)


			// If the price meets the target, update the notification status
			if price >= notification.TargetPrice {
				notification.Status = "Served"
				notification.ServedAt = time.Now().Format(time.RFC3339)
				color.New(color.FgGreen).Printf("Notification met: %s has reached $%.2f.\n", notification.Crypto, notification.TargetPrice)
			}
		}
	}

	// Save the updated notifications back to the file
	if err := savePriceNotifications(notifications); err != nil {
		color.New(color.FgRed).Printf("Error saving notifications: %v\n", err)
	}
}

// filterNotificationsByUsername filters notifications by username
func filterNotificationsByUsername(notifications []*models.PriceNotification, username string) []*models.PriceNotification {
	var filtered []*models.PriceNotification
	for _, notification := range notifications {
		if notification.Username == username {
			filtered = append(filtered, notification)
		}
	}
	return filtered
}

// Save the notifications to the JSON file
func savePriceNotifications(notifications []*models.PriceNotification) error {
	data, err := json.Marshal(notifications)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("price_notifications.json", data, 0644); err != nil {
		return err
	}

	return nil
}
