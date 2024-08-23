package auth

import (

	"cryptotracker/models"
	"cryptotracker/pkg/storage"
	"cryptotracker/pkg/utils"
	"errors"
	"fmt"
	"github.com/fatih/color"
)

// Login handles the login process
func Login() (*models.User, string, error) {
	var username, password string

	color.New(color.FgCyan).Print("Enter username: ")
	fmt.Scan(&username)
	password = utils.GetHiddenInput("Enter password: ")

	user, err := storage.GetUserByUsername(username)
	if err != nil {
		return nil, "", err
	}

	hashedPassword := utils.HashPassword(password)
	if user.Password != hashedPassword {
		return nil, "", errors.New("invalid username or password")
	}

	color.New(color.FgGreen).Println("Login successful.")


	return user, user.Role, nil
}
