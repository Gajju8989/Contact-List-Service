package mysql

import (
	"context"
	"fmt"
	"task2/database/mysql/model/auth_model"
)

func (i *impl) CreateUser(ctx context.Context, userdata *auth_model.UserRegistrations) error {
	existingUser := &auth_model.UserRegistrations{}
	err := i.mysqlDb.WithContext(ctx).Where("phone_number=?", userdata.PhoneNumber).First(existingUser).Error
	if err == nil {
		// A user with the provided phone number already exists
		return fmt.Errorf("user with phone number %s is already registered", userdata.PhoneNumber)
	}
	// If no user with the provided phone number exists, create the new user
	err = i.mysqlDb.WithContext(ctx).Create(userdata).Error
	if err != nil {
		return err
	}
	return nil
}
