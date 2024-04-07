package service

import (
	"context"
	"task2/contactapp/service/model"
	"task2/database/mysql/model/contact_list_model"
)

func (s *impl) UpdateSpamStatus(ctx context.Context, req *model.ContactListRequest) error {

	_, err := s.mySqlStore.GetAuthToken(ctx, req.Requester.RequesterPhoneNo)
	if err != nil {
		return err
	}
	createSqlPayload := &contact_list_model.GlobalContactList{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Spam:        "true",
	}
	err = s.mySqlStore.UpdateSpamStatus(ctx, createSqlPayload)

	if err != nil {
		return err
	}
	return nil
}
