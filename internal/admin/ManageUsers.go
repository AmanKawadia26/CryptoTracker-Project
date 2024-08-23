package admin

import (
	//"cryptotracker/models"
	"cryptotracker/pkg/storage"
	"fmt"
	"github.com/fatih/color"
)

func ManageUsers() {
	fmt.Println()
	color.New(color.FgGreen).Println("Managing users")
	fmt.Println("1. Change a user status to admin")
	fmt.Println("2. Delete a user")

	var choice int
	color.New(color.FgYellow).Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		ChangeUserStatus()
	case 2:
		DeleteUser()
	default:
		color.Red("Invalid choice")
	}

	fmt.Println()
}

func ChangeUserStatus() {
	// Load users from the storage package
	users, err := storage.LoadUsers()
	if err != nil {
		color.Red("Failed to load users: %v", err)
		return
	}

	// Get username input
	fmt.Print("Enter the username to change role: ")
	var username string
	fmt.Scan(&username)

	// Search for the user and update the role
	for _, user := range users {
		if user.Username == username {
			if user.Role == "admin" {
				color.Red("User is already an admin.")
				return
			}

			// Confirm role change
			fmt.Print("Change role to admin? (yes/no): ")
			var response string
			fmt.Scan(&response)

			if response == "yes" {
				user.Role = "admin"
				user.IsAdmin = true

				// Save the updated user
				if err := storage.SaveUser(user); err != nil {
					color.Red("Error saving user: %v", err)
				} else {
					color.Green("User role changed to admin successfully.")
				}
			}
			return
		}
	}

	color.Red("User not found.")
}

func DeleteUser() {
	// Load users from the storage package
	users, err := storage.LoadUsers()
	if err != nil {
		color.Red("Failed to load users: %v", err)
		return
	}

	// Get username input
	fmt.Print("Enter the username to delete: ")
	var username string
	fmt.Scan(&username)

	// Search for the user and delete them
	for i, user := range users {
		if user.Username == username {
			fmt.Print("Are you sure you want to delete this user? (yes/no): ")
			var response string
			fmt.Scan(&response)

			if response == "yes" {
				// Remove the user from the slice
				users = append(users[:i], users[i+1:]...)

				// Rewrite the whole user list
				if err := storage.SaveUsers(users); err != nil {
					color.Red("Error saving users: %v", err)
				} else {
					color.Green("User deleted successfully.")
				}
			}
			return
		}
	}

	color.Red("User not found.")
}
