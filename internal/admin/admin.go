package admin

import (
	"fmt"

	"github.com/fatih/color"
)

func ShowAdminPanel() {
	for {
		color.New(color.FgGreen).Println("Admin Panel")
		fmt.Println("1. Manage Users")
		fmt.Println("2. View User Profiles")
		fmt.Println("3. Manage User Requests")
		fmt.Println("4. Logout")

		var choice int
		color.New(color.FgYellow).Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ManageUsers()
		case 2:
			ViewUserProfiles()
		case 3:
			ManageUserRequests()
		case 4:
			color.New(color.FgCyan).Println("Logging out...")
			return
		default:
			color.New(color.FgRed).Println("Invalid choice, please try again.")
		}
	}
}
