package storage

import (
	"cryptotracker/models"
	"encoding/json"
	"errors"
	//"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

const userFile = "users.json"

// SaveUser saves a user to a file
func SaveUser(user *models.User) error {
	users, err := LoadUsers()
	if err != nil {
		color.Red("Failed to load users: %v", err)
		return err
	}

	users = append(users, user)
	data, err := json.Marshal(users)
	if err != nil {
		color.Red("Failed to marshal users: %v", err)
		return err
	}

	if err := ioutil.WriteFile(userFile, data, 0644); err != nil {
		color.Red("Failed to write user file: %v", err)
		return err
	}

	color.Green("User %s saved successfully!", user.Username)
	return nil
}

func SaveUsers(users []*models.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		color.Red("Failed to marshal users: %v", err)
		return err
	}

	if err := ioutil.WriteFile(userFile, data, 0644); err != nil {
		color.Red("Failed to write user file: %v", err)
		return err
	}

	color.Green("Users saved successfully!")
	return nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.User, error) {
	users, err := LoadUsers()
	if err != nil {
		color.Red("Failed to load users: %v", err)
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			color.Green("User %s found!", username)
			return user, nil
		}
	}

	color.Yellow("User %s not found.", username)
	return nil, errors.New("user not found")
}

// LoadUsers loads users from the file
func LoadUsers() ([]*models.User, error) {
	if _, err := os.Stat(userFile); os.IsNotExist(err) {
		color.Yellow("User file does not exist. Creating a new one.")
		return []*models.User{}, nil
	}

	data, err := ioutil.ReadFile(userFile)
	if err != nil {
		color.Red("Failed to read user file: %v", err)
		return nil, err
	}

	var users []*models.User
	if err := json.Unmarshal(data, &users); err != nil {
		color.Red("Failed to unmarshal users: %v", err)
		return nil, err
	}

	return users, nil
}

// GetUserProfile retrieves a user's profile by username
func GetUserProfile(username string) (*models.User, error) {
	users, err := LoadUsers()
	if err != nil {
		color.Red("Failed to load users: %v", err)
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			color.Green("User %s profile retrieved successfully!", username)
			return user, nil
		}
	}

	color.Yellow("Profile for user %s not found.", username)
	return nil, errors.New("user not found")
}
