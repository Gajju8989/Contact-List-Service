package mysql

import (
	"context"
	"gorm.io/gorm"
	"task2/database/mysql/model/auth_model"
	"task2/database/mysql/model/authtoken"
	"task2/database/mysql/model/contact_list_model"
)

type MySqlStore interface {
	//CreateUser
	CreateUser(ctx context.Context, userdata *auth_model.UserRegistrations) error
	//AuthToken
	StoreAuthToken(ctx context.Context, auth *authtoken.AuthToken) error
	GetAuthToken(ctx context.Context, phoneNo string) (*authtoken.AuthToken, error)

	//CreateContact
	StoreContact(ctx context.Context, createContactData *contact_list_model.ContactList) error
}

type impl struct {
	mysqlDb *gorm.DB
}

func NewMySqlStore(db *gorm.DB) MySqlStore {
	return &impl{mysqlDb: db}
}
