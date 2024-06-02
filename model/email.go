package model

type Email struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

type VerifyEmail struct {
	Link  string `json:"link"`
	Email string `json:"email"`
}
