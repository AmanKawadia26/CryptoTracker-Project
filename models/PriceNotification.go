package models

type PriceNotification struct {
	CryptoID    int     `json:"crypto_id"`
	Crypto      string  `json:"crypto"`              // The symbol of the cryptocurrency
	TargetPrice float64 `json:"target_price"`        // The price the user wants to track
	Username    string  `json:"username"`            // The username of the user who created the alert
	AskedAt     string  `json:"asked_at"`            // The timestamp when the alert was created
	Status      string  `json:"status"`              // The status of the alert (e.g., "Pending" or "Served")
	ServedAt    string  `json:"served_at,omitempty"` // The timestamp when the alert was triggered
}
