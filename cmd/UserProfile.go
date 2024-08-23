package main

import (
	"cryptotracker/pkg/storage"
	"fmt"
	"github.com/fatih/color"
)

func UserProfile(username string) {
	// Fetch user profile
	user, err := storage.GetUserProfile(username)
	if err != nil {
		color.New(color.FgRed).Println("Error fetching user profile:", err)
		return
	}

	// Display user profile
	fmt.Println()
	color.New(color.FgGreen).Println("User Profile:")
	color.New(color.FgCyan).Printf("Username: %s\n", user.Username)
	color.New(color.FgCyan).Printf("Email: %s\n", user.Email)
	color.New(color.FgCyan).Printf("Mobile: %d\n", user.Mobile)
	color.New(color.FgCyan).Printf("Role: %s\n", user.Role)
	fmt.Println()
}
