package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jacq/email"
	"jacq/encrypt"
	"jacq/generator"
	"jacq/helper"
	"jacq/idgenerator"
	"jacq/model"
	"jacq/storage"
	"net/http"
	"strings"
)

type UserHandler struct {
	store   storage.DataStore
	encrypt encrypt.Encryptor
	idGen   idgenerator.IdGenerator
	//Add logger here
}

func NewUserHandler(store storage.DataStore, enc encrypt.Encryptor, idGen idgenerator.IdGenerator) UserHandler {
	return UserHandler{
		store:   store,
		encrypt: enc,
		idGen:   idGen,
	}
}

func (user *UserHandler) CreateUser(data *model.User) error {
	//Prevent duplicate email
	_, err := user.store.GetUserByEmail(data.Email)
	if err == nil {
		err := fmt.Errorf("error user email already exist %v", err)
		return err
	}

	//Check if the password and confirm password matches
	if data.Password != data.ConfirmPassword {
		err := fmt.Errorf("error password and confirm password are not matched %v", err)
		return err
	}

	//Hash the user password
	encryptedPassword, err := user.encrypt.HashPassword(data.Password)
	if err != nil {
		err := fmt.Errorf("error encrypting password %v", err)
		return err
	}

	//Trim the email
	trimmedEmail := strings.TrimSpace(data.Email)

	d := &model.User{
		ID:              user.idGen.Generate(),
		Email:           trimmedEmail,
		Password:        encryptedPassword,
		ConfirmPassword: encryptedPassword,
	}

	//Persist the data to the db...
	_, err = user.store.CreateUser(d)
	if err != nil {
		err := fmt.Errorf("error creating a user %v", err)
		return err
	}
	return nil
}

func (user *UserHandler) SendVerificationLink(userEmail, link string) error {
	//Ensure the user's record exist
	_, err := user.store.GetUserByEmail(strings.TrimSpace(userEmail))
	if err != nil {
		err := fmt.Errorf("error user's record not exist %v", err)
		return err
	}
	//Generate OTP for verification
	otp, err := helper.GenerateOTP()
	if err != nil {
		err := fmt.Errorf("error generating OTP for verification %v", err)
		return err
	}

	//Prepare email body
	emailBody := link + otp
	data := model.Email{
		To:   userEmail,
		Body: emailBody,
	}

	return email.SendEmailVerification(data)
}

func (user *UserHandler) UpdateUser(ID string, data *model.User) error {
	//Check if the user's record exists
	_, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error user's record is not exist %v", err)
		return err
	}
	//Trim user's data
	firstName := strings.TrimSpace(data.FirstName)
	lastName := strings.TrimSpace(data.LastName)
	address := strings.TrimSpace(data.HomeAddress)

	record := &model.User{
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: data.DateOfBirth,
		Gender:      data.Gender,
		Country:     data.Country,
		State:       data.State,
		HomeAddress: address,
	}

	//Update the user record
	_, err = user.store.UpdateUser(ID, record)
	if err != nil {
		err := fmt.Errorf("error updating user's record %v", err)
		return err
	}
	return nil
}

func (user *UserHandler) VerifyNumber(ID, phoneNumber string) error {
	//Ensure the user's record exist
	_, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error updating user's record %v", err)
		return err
	}

	//Generate an OTP for a user
	code, err := helper.GenerateOTP()
	if err != nil {
		err := fmt.Errorf("error generating OTP for a user %v", err)
		return err
	}

	//Prepare an SMS
	rec := model.VerifyPhoneNumber{
		Body: code,
		From: phoneNumber, //We need to buy a phone number
	}

	return helper.VerifyNumber(rec)
}

func (user *UserHandler) VerifyBvn(ID string, data *model.User) error {
	//Ensure that user's record is available
	_, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error user's record not found %v", err)
		return err
	}
	//https://vapi.verifyme.ng/v1/verifications/identities/bvn/10000000001 This is url from VerifyMe doc
	method := "POST"
	payload := map[string]string{
		"bvn": data.Bvn,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		err := fmt.Errorf("error marshaling payload %v", err)
		return err
	}

	body := bytes.NewBuffer(payloadBytes)
	//Initialize the http request
	client := &http.Client{}
	req, err := http.NewRequest(method, helper.Url, body)
	if err != nil {
		err := fmt.Errorf("error creating http request%v", err)
		return err
	}

	//Send the http request
	res, err := client.Do(req)
	if err != nil {
		err := fmt.Errorf("error sending http request %v", err)
		return err
	}
	defer res.Body.Close()

	//Validate if the http was successful
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("error http request was unsuccessful %v", err)
		return err
	}

	//Read the response http response body
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err := fmt.Errorf("error reading http response body %v", err)
		return err
	}

	//Serialize the data from the response body
	var verifyMeResponse model.VerifyMeResponse
	err = json.Unmarshal(bodyBytes, &verifyMeResponse)
	if err != nil {
		err := fmt.Errorf("error unmarshaling http response %v", err)
		return err
	}

	if verifyMeResponse.Status != "success" {
		err := fmt.Errorf("error bvn verification failed %v", err)
		return err
	}

	//If bvn validation is successful then we update the user's record
	d := &model.User{
		IDType:   data.IDType,
		Document: data.Document,
	}
	_, err = user.store.UpdateUser(ID, d)
	if err != nil {
		err := fmt.Errorf("error updating user's record %v", err)
		return err
	}
	return nil
}

func (user *UserHandler) SecureTransaction(ID string, data *model.User) error {
	//Verify if the user's record exist
	_, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error user`s record is not exist %v", err)
		return err
	}
	fmt.Println(data)

	//Secure the user transaction code
	record := &model.User{
		TransactionCode: data.TransactionCode,
	}
	_, err = user.store.UpdateUser(ID, record)
	if err != nil {
		err := fmt.Errorf("error updating user's record %v", err)
		return err
	}
	return nil
}

func (user *UserHandler) Login(data *model.User) (string, error) {
	//Ensure that the user's record exist
	us, err := user.store.GetUserByEmail(data.Email)
	if err != nil {
		err := fmt.Errorf("error user's record not exist %v", err)
		return "", err
	}
	//Compare the user's passwords
	_, err = user.encrypt.CompareHashAndPassword(us.Password, data.Password)
	if err != nil {
		err := fmt.Errorf("error user's passwords are not matched %v", err)
		return "", err
	}

	//Generate access token
	token, err := generator.GenerateAccessToken(us.ID, us.Email)
	if err != nil {
		err := fmt.Errorf("error generating user's access token %v", err)
		return "", err
	}
	return token, nil
}

func (user *UserHandler) ForgotPassword(data *model.User) error {
	//Ensure user's record exist
	trimedEmail := strings.TrimSpace(data.Email)

	//Ensure user's record exist
	_, err := user.store.GetUserByEmail(trimedEmail)
	if err != nil {
		err := fmt.Errorf("error user's record not exist %v", err)
		return err
	}

	//Generate an OTP
	otp, err := helper.GenerateOTP()
	if err != nil {
		err := fmt.Errorf("error generating OTP %v", err)
		return err
	}
	//Create a link
	var link = "Pls provide me a real link"
	var body = link + otp

	record := model.Email{
		To:   trimedEmail,
		Body: body,
	}
	return email.SendEmailVerification(record)
}

func (user *UserHandler) ChangePassword(ID string, data *model.ChangePassword) error {
	//Ensure user's record exist
	us, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error getting user's record %v", err)
		return err
	}

	//Compare current password with the existing password
	_, err = user.encrypt.CompareHashAndPassword(us.Password, data.CurrentPassword)
	if err != nil {
		err := fmt.Errorf("error comparing current password with the existing one %v", err)
		return err
	}

	//Compare New password with confirm password
	if data.NewPassword != data.ConfirmPassword {
		err := fmt.Errorf("error comparing new password with the confirm password  %v", err)
		return err
	}

	//Encrypt the new password
	encrPass, err := user.encrypt.HashPassword(data.NewPassword)
	if err != nil {
		err := fmt.Errorf("error encrypting new password  %v", err)
		return err
	}

	//Update user's password with new password
	rec := &model.User{
		Password:        encrPass,
		ConfirmPassword: encrPass,
	}
	_, err = user.store.UpdateUser(ID, rec)
	if err != nil {
		err := fmt.Errorf("error updating user's record %v", err)
		return err
	}

	//Send email to the user
	var body = "Password changed successfully"

	record := model.Email{
		To:   us.Email,
		Body: body,
	}
	return email.SendPasswordChangedSuccessfully(record)
}

func (user *UserHandler) ChangeTransactionPin(ID string, data *model.TransactionPin) error {
	//Ensure that user's record exist
	us, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error user's record not exist %v", err)
		return err
	}
	//Compare user current pin with the user database existing pin
	if us.TransactionCode != data.CurrentPin {
		err := fmt.Errorf("error user's record pin does match current pin supplied %v", nil)
		return err
	}

	//Ensure that both newPin and ConfirmNewPin are not empty
	if data.NewPin == 0 || data.ConfirmNewPin == 0 {
		err := fmt.Errorf("error user's newPin and ConfirmNewPin are empty %v", nil)
		return err
	}

	//Ensure both the newPin and confirmNewPin are the same
	if data.NewPin != data.NewPin {
		err := fmt.Errorf("error user's newPin and ConfirmNewPin are not matched %v", nil)
		return err
	}

	//Ensure that newPin is not the same as currentPin
	if data.NewPin == data.CurrentPin {
		err := fmt.Errorf("error user's newPin and CurrentPin are not same %v", nil)
		return err
	}

	//Update the user's transaction pin
	rec := &model.User{
		TransactionCode: data.NewPin,
	}
	_, err = user.store.UpdateUser(ID, rec)
	if err != nil {
		err := fmt.Errorf("error updating user transaction pin %v", nil)
		return err
	}

	//Send user an email for pin successfully changed
	d := model.Email{
		To:   us.Email,
		Body: "Pin changed successfully",
	}
	return email.SendPinChangedSuccessfully(d)
}

func (user *UserHandler) VerifyOtp(ID, authToken string) error {
	//Ensure user's record is available
	_, err := user.store.GetUserByID(ID)
	if err != nil {
		err := fmt.Errorf("error user's record is not available %v", nil)
		return err
	}
	//Verify the OTP
	return generator.ValidateAccessToken(authToken)
}
