package main

import (
	"cryptotracker/pkg/config"
	"cryptotracker/pkg/ui"
)

func main() {
	// Load the configuration
	config.LoadConfig()

	// Display welcome banner
	ui.DisplayWelcomeBanner()

	// Start login/signup process
	user, Role := ui.AuthenticateUser()

	// If user is admin, show admin panel
	if Role == "admin" {
		ui.ShowAdminPanel()
		return
	}

	// Display main user menu
	ui.MainMenu(user)
}
