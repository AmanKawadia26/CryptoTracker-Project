package models

type Request struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Symbol    string `json:"symbol"`
	Status    string `json:"status"` // Pending, Approved, Disapproved
	DateAdded string `json:"date_added"`
}
