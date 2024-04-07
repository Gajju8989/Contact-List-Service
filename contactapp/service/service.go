package service

import (
	"context"
	"task2/contactapp/service/model"
	"task2/database/mysql"
)

type ContactAppSvc interface {
	CreateUser(ctx context.Context, req *model.UserRegistrationsRequestModel) error

	CreateContact(ctx context.Context, req *model.ContactListRequest) error
}

type impl struct {
	mySqlStore mysql.MySqlStore
}

func NewUserSvc(mySqlStore mysql.MySqlStore) ContactAppSvc {
	return &impl{mySqlStore: mySqlStore}
}
