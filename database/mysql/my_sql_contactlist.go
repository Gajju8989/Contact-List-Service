package mysql

import (
	"context"
	"task2/database/mysql/model/contact_list_model"
)

func (m *impl) StoreContact(ctx context.Context, createContactData *contact_list_model.ContactList) error {
	return m.mysqlDb.
		WithContext(ctx).
		Create(&createContactData).
		Error
}
