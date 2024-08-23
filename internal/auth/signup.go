package auth

import (
	"cryptotracker/models"
	"cryptotracker/pkg/storage"
	"cryptotracker/pkg/utils"
	"cryptotracker/pkg/validation"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Signup handles the signup process
func Signup() error {

	users, err := storage.LoadUsers()
	if err != nil {
		log.Fatal(err)
	}

	var username, password, email, pan string
	var mobile int

	color.New(color.FgCyan).Print("Enter username: ")
	fmt.Scan(&username)
	if !validation.IsValidUsername(username) {
		return errors.New("invalid username: must be one word, alphanumeric, and can contain underscores")
	}
	for _, user := range users {
		if user.Username == username {
			return errors.New("user already exists")
		}
	}

	password = utils.GetHiddenInput("Enter password: ")
	if !validation.IsValidPassword(password) {
		return errors.New("invalid password: must be at least 8 characters, include an uppercase letter, a number, and a special character")
	}

	color.New(color.FgCyan).Print("Enter email: ")
	fmt.Scan(&email)
	if !validation.IsValidEmail(email) {
		return errors.New("invalid email: must be a valid email address")
	}

	color.New(color.FgCyan).Print("Enter mobile (10 digits): ")
	fmt.Scan(&mobile)
	if !validation.IsValidMobile(mobile) {
		return errors.New("invalid mobile number: must be 10 digits")
	}

	color.New(color.FgCyan).Print("Enter PAN: ")
	fmt.Scan(&pan)

	hashedPassword := utils.HashPassword(password)

	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Mobile:   mobile,
		IsAdmin:  false,
		Role:     "user",
	}

	if err := storage.SaveUser(user); err != nil {
		return err
	}

	color.New(color.FgGreen).Println("Signup successful.")
	return nil
}
