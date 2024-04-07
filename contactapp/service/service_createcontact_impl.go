package service

import (
	"context"
	"task2/contactapp/service/model"
	"task2/database/mysql/model/contact_list_model"
)

func (s *impl) CreateContact(ctx context.Context, req *model.ContactListRequest) error {

	//Checking If Token Exist For RequesterPhoneNo or Token is expired
	_, err := s.mySqlStore.GetAuthToken(ctx, req.Requester.RequesterPhoneNo)
	if err != nil {
		return err
	}
	//Creating an Entry to store contactDetails in GlobalContactList Table MySql.
	createContactData := &contact_list_model.GlobalContactList{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	//StoringInDb
	err = s.mySqlStore.StoreContactListInGlobalDb(ctx, createContactData)
	if err != nil {
		return err
	}
	return nil
}
