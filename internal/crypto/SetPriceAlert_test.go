package crypto

import "testing"

func Test_SetPriceAlert(t *testing.T) {
	tests := []struct {
		name           string
		apiResponse    []byte
		userInput      string
		expectedOutput string
	}{
		{
			name: "Alert Triggered",
			apiResponse: []byte(`{
				"data": {
					"BTC": {
						"quote": {
							"USD": {
								"price": 50000
							}
						}
					}
				}
			}`),
			userInput:      "BTC\n49000\n",
			expectedOutput: "Alert: BTC has reached your target price of $49000.00. Current price: $50000.00\n",
		},
		{
			name: "No Alert",
			apiResponse: []byte(`{
				"data": {
					"BTC": {
						"quote": {
							"USD": {
								"price": 49000
							}
						}
					}
				}
			}`),
			userInput:      "BTC\n50000\n",
			expectedOutput: "BTC is still below your target price. Current price: $49000.00. Notification created.\n",
		},
		{
			name: "Cryptocurrency Not Found",
			apiResponse: []byte(`{
				"data": {}
			}`),
			userInput:      "BTC\n50000\n",
			expectedOutput: "Cryptocurrency data not found for symbol: BTC\n",
		},
	}

	for _, tt := range tests {

	}

}
