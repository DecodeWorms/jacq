package model

// VerifyMeResponse represents the response structure from VerifyMe API
type VerifyMeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		DOB       string `json:"dob"`
	} `json:"data"`
}
