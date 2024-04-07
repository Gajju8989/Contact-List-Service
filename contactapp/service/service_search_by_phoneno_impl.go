package service

import (
	"context"
	"task2/contactapp/service/model"
	"task2/database/mysql/model/contact_list_model"
)

func (s *impl) SearchUserByPhoneNo(ctx context.Context, phoneNo string, requesterPhoneNo string) ([]*model.ContactListResponseModel, error) {
	_, err := s.mySqlStore.GetAuthToken(ctx, requesterPhoneNo)
	if err != nil {
		return nil, err
	}
	var res []*contact_list_model.GlobalContactList
	res, err = s.mySqlStore.SearchByPhoneNumber(ctx, phoneNo)
	if err != nil {
		return nil, err
	}
	var ContactListData []*model.ContactListResponseModel
	for _, contact := range res {
		contactData := &model.ContactListResponseModel{
			Name:        contact.Name,
			PhoneNumber: contact.PhoneNumber,
			Spam:        contact.Spam,
		}
		ContactListData = append(ContactListData, contactData)
	}

	return ContactListData, nil
}
