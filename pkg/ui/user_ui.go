package ui

import (
	"cryptotracker/internal/crypto"
	user2 "cryptotracker/internal/user"
	"cryptotracker/models"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

// MainMenu displays the main menu for a regular user
func MainMenu(user *models.User) {
	for {
		ClearScreen()
		DisplayMainMenu()

		var choice int
		color.New(color.FgYellow).Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			crypto.DisplayTopCryptocurrencies()
		case 2:
			crypto.SearchCryptocurrency()
		case 3:
			crypto.SetPriceAlert(user)
		case 4:
			user2.UserProfile(user.Username)
		case 5:
			color.New(color.FgCyan).Println("Logging out...")
			log.Println("Logging out...")
			os.Exit(0)
		default:
			color.New(color.FgRed).Println("Invalid choice, please try again.")
		}
	}
}
