package models

type UnavailableCryptoRequest struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	CryptoSymbol   string `json:"crypto_symbol"`
	RequestMessage string `json:"request_message"`
	Status         string `json:"status"` // Possible values: Pending, Approved, Rejected
}
