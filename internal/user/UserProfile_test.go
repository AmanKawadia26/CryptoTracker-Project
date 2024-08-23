package user

import (
	"bytes"
	"cryptotracker/models"
	"cryptotracker/pkg/storage"
	"errors"
	"fmt"
	"testing"

	"github.com/fatih/color"
)

// MockStorage simulates the storage package behavior for testing
type MockStorage struct {
	user *models.User
	err  error
}

// Mock GetUserProfile function
func (m *MockStorage) GetUserProfile(username string) (*models.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.user, nil
}

// Replace the original GetUserProfile call in the user package with this mock implementation
var mockStorage *MockStorage

func init() {
	storage.GetUserProfile = mockStorage.GetUserProfile
}

// Helper function to capture the output of functions that print to the console
func captureOutput(f func()) string {
	var buf bytes.Buffer
	original := color.Output
	color.Output = &buf
	defer func() { color.Output = original }()

	f()

	return buf.String()
}

func TestUserProfileSuccess(t *testing.T) {
	// Set up mock storage with a valid user profile
	mockStorage = &MockStorage{
		user: &models.User{
			Username: "testuser",
			Email:    "test@example.com",
			Mobile:   1234567890,
			Role:     "user",
		},
	}

	// Capture the output
	output := captureOutput(func() {
		UserProfile("testuser")
	})

	// Verify the output contains the expected user profile details
	expectedOutput := fmt.Sprintf("%s\n%s%s\n%s%s\n%s%d\n%s%s\n",
		color.GreenString("User Profile:"),
		color.CyanString("Username: "), "testuser",
		color.CyanString("Email: "), "test@example.com",
		color.CyanString("Mobile: "), 1234567890,
		color.CyanString("Role: "), "user",
	)

	if output != expectedOutput {
		t.Errorf("Expected output:\n%q\nbut got:\n%q", expectedOutput, output)
	}
}

func TestUserProfileError(t *testing.T) {
	// Set up mock storage to return an error
	mockStorage = &MockStorage{
		err: errors.New("user not found"),
	}

	// Capture the output
	output := captureOutput(func() {
		UserProfile("nonexistentuser")
	})

	// Verify the error message is printed
	expectedOutput := color.RedString("Error fetching user profile: user not found\n")

	if output != expectedOutput {
		t.Errorf("Expected output:\n%q\nbut got:\n%q", expectedOutput, output)
	}
}
