package admin

import (
	"cryptotracker/pkg/storage"
	"github.com/fatih/color"
)

func ViewUserProfiles() {
	users, err := storage.GetAllUsers()
	if err != nil {
		color.New(color.FgRed).Println("Error fetching user profiles:", err)
		return
	}

	if len(users) == 0 {
		color.New(color.FgCyan).Println("No users found.")
		return
	}

	for i, user := range users {
		if user.Role == "user" {
			color.New(color.FgYellow).Printf("[%d] - Username: %s, Email: %s, Mobile: %d, Role: %s\n", i+1, user.Username, user.Email, user.Mobile, user.Role)
		}
	}
}
