package model

type VerifyPhoneNumber struct {
	Body string `json:"body"`
	From string `json:"from"`
	To   string `json:"to"`
	Code string `json:"code"`
}
