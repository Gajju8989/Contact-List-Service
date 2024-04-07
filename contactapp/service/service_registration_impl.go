package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"task2/contactapp/service/model"
	"task2/database/mysql/model/auth_model"
	"task2/database/mysql/model/authtoken"
	"time"
)

func (s *impl) CreateUser(ctx context.Context, userData *model.UserRegistrationsRequestModel) error {

	// Creating the UserRegistrations object to save in the database
	saveUserData := &auth_model.UserRegistrations{
		Name:        userData.Name,
		PhoneNumber: userData.PhoneNumber,
		Email:       userData.Email,
		Password:    userData.Password,
	}
	//Saving The RegistrationData in Mysql
	err := s.mySqlStore.CreateUser(ctx, saveUserData)
	if err != nil {
		return err
	}
	token, err := generateJWTToken(userData.PhoneNumber)
	if err != nil {
		return err
	}
	createAuthTokenData := &authtoken.AuthToken{
		PhoneNo: userData.PhoneNumber,
		Token:   token,
		Expiry:  time.Now().Add(24 * 30 * time.Hour).Unix(),
	}
	err = s.mySqlStore.StoreAuthToken(ctx, createAuthTokenData)
	if err != nil {
		return err
	}
	return nil
}

func generateJWTToken(phoneNumber string) (string, error) {

	expirationTime := time.Now().Add(24 * 30 * time.Hour) // Token expires in 1 Month
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   phoneNumber, // Use PhoneNumber as subject
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key")) // Use a secure secret key
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
