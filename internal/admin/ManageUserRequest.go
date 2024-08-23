package admin

import (
	"cryptotracker/pkg/storage"
	"fmt"
	"github.com/fatih/color"
)

func ManageUserRequests() {
	// Add handling for unavailable crypto requests
	unavailableRequests, err := storage.GetAllUnavailableCryptoRequests()
	if err != nil {
		color.New(color.FgRed).Println("Error fetching unavailable crypto requests:", err)
		return
	}

	if len(unavailableRequests) == 0 {
		color.New(color.FgCyan).Println("No pending unavailable crypto requests.")
		return
	}

	for i, request := range unavailableRequests {
		color.New(color.FgYellow).Printf("[%d] - Symbol: %s, User: %s, Message: %s, Status: %s\n", i+1, request.CryptoSymbol, request.UserID, request.RequestMessage, request.Status)
	}

	var choice int
	color.New(color.FgYellow).Print("Select a request to manage: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(unavailableRequests) {
		color.New(color.FgRed).Println("Invalid selection.")
		return
	}

	selectedRequest := unavailableRequests[choice-1]
	color.New(color.FgGreen).Println("Selected request:", selectedRequest)

	// Allow admin to approve or reject the request
	var action string
	color.New(color.FgYellow).Print("Enter 'approve' to approve the request or 'reject' to reject it: ")
	fmt.Scan(&action)

	if action == "approve" {
		selectedRequest.Status = "Approved"
	} else if action == "reject" {
		selectedRequest.Status = "Rejected"
	} else {
		color.New(color.FgRed).Println("Invalid action.")
		return
	}

	if err := storage.SaveUnavailableCryptoRequest(selectedRequest); err != nil {
		color.New(color.FgRed).Println("Error updating request status:", err)
	} else {
		color.New(color.FgGreen).Println("Request status updated.")
	}
}
