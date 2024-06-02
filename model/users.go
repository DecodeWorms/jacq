package model

import "time"

type User struct {
	ID              string    `json:"id" bson:"id"`
	Email           string    `json:"email" bson:"email"`
	Password        string    `json:"password" bson:"password"`
	ConfirmPassword string    `json:"confirm_password" bson:"confirm_password"`
	FirstName       string    `json:"first_name" bson:"first_name"`
	LastName        string    `json:"last_name" bson:"last_name"`
	DateOfBirth     string    `json:"date_of_birth" bson:"date_of_birth"`
	Gender          string    `json:"gender" bson:"gender"`
	Country         string    `json:"country" bson:"country"`
	State           string    `json:"state" bson:"state"`
	HomeAddress     string    `json:"home_address" bson:"home_address"`
	PhoneNumber     string    `json:"phone_number" bson:"phone_number"`
	Bvn             string    `json:"bvn" bson:"bvn"`
	IDType          string    `json:"IDType" bson:"IDType"`
	TransactionCode int       `json:"transaction_code" bson:"transaction_code"`
	StatusTimeStamp time.Time `json:"status_ts" bson:"status_ts"`
	Timestamp       time.Time `json:"ts" bson:"ts"`
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
