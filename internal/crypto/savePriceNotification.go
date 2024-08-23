package crypto

import (
	"cryptotracker/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Save a single notification to the JSON file
func savePriceNotification(notification *models.PriceNotification) error {
	notifications, err := LoadPriceNotifications()
	if err != nil {
		return err
	}

	notifications = append(notifications, notification)
	data, err := json.Marshal(notifications)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("price_notifications.json", data, 0644); err != nil {
		return err
	}

	return nil
}

func LoadPriceNotifications() ([]*models.PriceNotification, error) {
	if _, err := os.Stat("price_notifications.json"); os.IsNotExist(err) {
		return []*models.PriceNotification{}, nil
	}

	data, err := ioutil.ReadFile("price_notifications.json")
	if err != nil {
		return nil, err
	}

	var notifications []*models.PriceNotification
	if err := json.Unmarshal(data, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}
