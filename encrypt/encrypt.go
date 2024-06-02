package encrypt

import "golang.org/x/crypto/bcrypt"

//go:generate mockgen -source=encrypt.go -destination=../../mocks/encrypt_mock.go -package=mocks
type Encryptor interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) (bool, error)
}

type passwordEncryptor struct{}

func NewPasswordEncryptor() Encryptor {
	return &passwordEncryptor{}
}

func (e *passwordEncryptor) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (e *passwordEncryptor) CompareHashAndPassword(hashedPassword, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
