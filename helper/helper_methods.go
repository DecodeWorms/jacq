package helper

import (
	"crypto/rand"
	"fmt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"jacq/config"
	"jacq/model"
	"math/big"
)

const (
	PhoneNumber = "BUY_ME_PHONE_NUMBER"
	Url         = "https://vapi.verifyme.ng/v1/verifications/bvn"
)

// GenerateOTP generates a 6-digit one-time password
func GenerateOTP() (string, error) {
	// Define the range for the OTP
	// const otpLength = 6
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

func VerifyNumber(data model.VerifyPhoneNumber) error {
	c := config.ImportConfig(config.OSSource{})
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.AccountSid,
		Password: c.AuthToken,
	})

	params := &api.CreateMessageParams{}
	params.SetBody(data.Body)
	params.SetFrom(PhoneNumber)
	params.SetTo(data.To)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		err := fmt.Errorf("error sending message to a phone number %v", err)
		return err
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
	return nil
}
