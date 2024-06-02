package handler

import (
	"fmt"
	"jacq/email"
	"jacq/encrypt"
	"jacq/generator"
	"jacq/helper"
	"jacq/idgenerator"
	"jacq/model"
	"jacq/storage"
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

	if err := email.SendEmail(data); err != nil {
		err := fmt.Errorf("error sending email %v", err)
		return err
	}
	return nil
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
	if err := helper.VerifyNumber(rec); err != nil {
		err := fmt.Errorf("error verifying phone number %v", err)
		return err
	}
	return nil
}

func (user *UserHandler) VerifyBvn(data *model.User) error {
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
	if err := email.SendEmail(record); err != nil {
		err := fmt.Errorf("error sending an email %v", err)
		return err
	}
	return nil
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
	return nil
}
