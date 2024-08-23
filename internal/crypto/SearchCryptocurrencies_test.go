package crypto

import (
	"bytes"
	"cryptotracker/internal/api"
	"fmt"
	"github.com/fatih/color"
	"strings"
	"testing"
)

// Mock API response function
func mockGetAPIResponse(endpoint string, params map[string]string) []byte {
	if endpoint == "/listings/latest" {
		return []byte(`{
			"data": [
				{
					"symbol": "btc",
					"name": "Bitcoin",
					"quote": {
						"USD": {
							"price": 50000
						}
					}
				}
			]
		}`)
	}
	return nil
}

// Mock function to simulate input and output
func mockInput(prompt string) string {
	return "btc"
}

func TestSearchCryptocurrency(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   []byte
		userInput      string
		expectedOutput string
	}{
		{
			name: "Cryptocurrency Found",
			mockResponse: []byte(`{
				"data": [
					{
						"symbol": "btc",
						"name": "Bitcoin",
						"quote": {
							"USD": {
								"price": 50000
							}
						}
					}
				]
			}`),
			userInput:      "btc",
			expectedOutput: "Bitcoin (BTC): $50000.00\n\n30-day price graph for Bitcoin:\n\n",
		},
		{
			name: "Cryptocurrency Not Found",
			mockResponse: []byte(`{
				"data": [
					{
						"symbol": "eth",
						"name": "Ethereum",
						"quote": {
							"USD": {
								"price": 3000
							}
						}
					}
				]
			}`),
			userInput:      "btc",
			expectedOutput: "Cryptocurrency not found for input: btc\nPlease request the addition of this cryptocurrency to our app.\nRequest to add the cryptocurrency has been submitted.\n",
		},
		{
			name: "API Response Error",
			mockResponse: []byte(`{
				"invalid": "response"
			}`),
			userInput:      "btc",
			expectedOutput: "Error unmarshalling API response: json: cannot unmarshal object into Go value of type []interface {}\n",
		},
	}

	originalGetAPIResponse := api.GetAPIResponse
	defer func() { api.GetAPIResponse = originalGetAPIResponse }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect output to buffer
			var buf bytes.Buffer
			color.Output = &buf

			// Override getAPIResponse with mock response
			api.GetAPIResponse = func(endpoint string, params map[string]string) []byte {
				return tt.mockResponse
			}

			// Temporarily replace fmt.Scan with a mock
			oldScan := fmt.Scan
			defer func() { fmt.Scan = oldScan }()
			fmt.Scan = func(a ...interface{}) (n int, err error) {
				switch v := a[0].(type) {
				case *string:
					*v = tt.userInput
				}
				return 0, nil
			}

			// Call the function under test
			SearchCryptocurrency()

			// Check the output
			if got := buf.String(); !strings.Contains(got, tt.expectedOutput) {
				t.Errorf("expected %q, got %q", tt.expectedOutput, got)
			}
		})
	}
}
