package model

type TransactionPin struct {
	CurrentPin    int `json:"current_pin"`
	NewPin        int `json:"new_pin"`
	ConfirmNewPin int `json:"confirm_new_pin"`
}
