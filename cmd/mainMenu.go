package main

import (
	"cryptotracker/internal/admin"
	"cryptotracker/internal/crypto"
	"cryptotracker/models"
	"cryptotracker/pkg/ui"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

// mainMenu displays the main menu for a regular user
func mainMenu(user *models.User) {
	for {
		ui.ClearScreen()
		ui.DisplayMainMenu()

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
			if user.Role == "admin" {
				admin.ShowAdminPanel()
			} else {
				color.New(color.FgRed).Println("This user is not an admin.\n")
			}
		case 5:
			UserProfile(user.Username)
		case 6:
			color.New(color.FgCyan).Println("Logging out...")
			log.Println("Logging out...")
			os.Exit(0)
		default:
			color.New(color.FgRed).Println("Invalid choice, please try again.")
		}
	}
}
