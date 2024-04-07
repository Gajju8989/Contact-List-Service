package service

import (
	"context"
	"fmt"
	"task2/contactapp/service/model"
	"task2/database/mysql/model/contact_list_model"
)

func (s *impl) CreateContact(ctx context.Context, req *model.ContactListRequest) error {

	res, err := s.mySqlStore.GetAuthToken(ctx, req.Requester.RequesterPhoneNo)
	if err != nil {
		return err
	}
	fmt.Println(res)
	fmt.Println(req.Requester.RequesterPhoneNo)
	createContactData := &contact_list_model.ContactList{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	err = s.mySqlStore.StoreContact(ctx, createContactData)
	if err != nil {
		return err
	}
	return nil
}
