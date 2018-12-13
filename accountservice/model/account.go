package model

// Account struct
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	ServedBy string `json:"servedBy"`
}
