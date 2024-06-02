package generator

import (
	"fmt"
	//"github.com/dgrijalva/jwt-go"

	//jwtn "github.com/dgrijalva/jwt-go"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateAccessToken(uId, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": uId,
		"exp":     time.Now().UTC().Add(time.Minute * 60).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(uId, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": uId,
		"exp":     time.Now().UTC().Add(time.Minute * 180).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateAccessToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return err
	}
	return nil
}
