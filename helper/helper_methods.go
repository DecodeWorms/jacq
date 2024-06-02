package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOTP generates a 6-digit one-time password
func GenerateOTP() (string, error) {
	// Define the range for the OTP
	const otpLength = 6
	const max = 1000000

	// Generate a secure random number within the range
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return "", err
	}

	// Format the number to ensure it is 6 digits
	otp := fmt.Sprintf("%06d", n.Int64())

	return otp, nil
}
