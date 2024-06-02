package model

type Will struct {
	ID           string `bson:"id" bson:"id"`
	UserID       string `json:"user_id" bson:"user_id"`
	WitnessName  string `json:"witness_name" bson:"witness_name"`
	WitnessEmail string `json:"witness_email" bson:"witness_email"`
	MeetingDate  string `json:"meeting_date" json:"meeting_date"`
	MeetingTime  string `json:"meeting_time" bson:"meeting_time"`
}

type Executor struct {
	ID          string `json:"id" bson:"id"`
	UserID      string `json:"user_id" bson:"user_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Address     string `json:"address" bson:"address"`
	State       string `json:"state" bson:"state"`
	Email       string `json:"email" bson:"email"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
}

type Beneficiary struct {
	ID          string `json:"id" bson:"id"`
	UserID      string `json:"user_id" bson:"user_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Address     string `json:"address" bson:"address"`
	State       string `json:"state" bson:"state"`
	Email       string `json:"email" bson:"email"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
}

type Assets struct {
	ID               string   `json:"id" bson:"id"`
	UserID           string   `json:"user_id" bson:"user_id"`
	AssetType        string   `json:"asset_type" bson:"asset_type"`
	PropertyDocument string   `json:"property_document" bson:"property_document"` //This is an image
	Beneficiary      []string `json:"beneficiary" bson:"beneficiary"`
}

type Witness struct {
	ID          string `json:"id" bson:"id"`
	UserID      string `json:"user_id" bson:"user_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Address     string `json:"address" bson:"address"`
	State       string `json:"state" bson:"state"`
	Country     string `json:"country" bson:"country"`
	Email       string `json:"email" bson:"email"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
}
