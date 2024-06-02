package storage

import (
	"jacq/model"
)

type DataStore interface {
	User
}

type User interface {
	CreateUser(data *model.User) (*model.User, error)
	UpdateUser(ID string, data *model.User) (*model.User, error)
	VerifyNumber(data *model.User) error
	VerifyIdentity(data *model.User) (*model.User, error)
	SecureTransaction(data *model.User) error
	GetUserByID(ID string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
