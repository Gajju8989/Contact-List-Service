package mysql

import (
	"context"
	"task2/database/mysql/model/contact_list_model"
)

func (m *impl) StoreContactListInGlobalDb(ctx context.Context, createContactData *contact_list_model.GlobalContactList) error {
	return m.mysqlDb.
		WithContext(ctx).
		Create(&createContactData).
		Error
}

func (m *impl) UpdateSpamStatus(ctx context.Context, createContactData *contact_list_model.GlobalContactList) error {
	// Define the fields to be updated
	updateFields := map[string]interface{}{
		"spam": createContactData.Spam,
	}

	// Update the database record if it exists
	result := m.mysqlDb.
		WithContext(ctx).
		Model(&contact_list_model.GlobalContactList{}).
		Where("phone_number = ?", createContactData.PhoneNumber).
		Updates(updateFields)

	// Check for errors
	if result.Error != nil {
		return result.Error
	}

	// Check if the record was found and updated
	if result.RowsAffected == 0 {
		// Phone number doesn't exist in the database, treat it as a new entry
		return m.mysqlDb.
			WithContext(ctx).
			Create(createContactData).Error
	}

	return nil
}

// SearchByName In MySql
func (m *impl) SearchByName(ctx context.Context, name string) ([]*contact_list_model.GlobalContactList, error) {
	var results []*contact_list_model.GlobalContactList

	query := m.mysqlDb.WithContext(ctx).Where("name LIKE ?", name+"%")

	query = query.Order("CASE WHEN name LIKE '" + name + "%' THEN 0 ELSE 1 END, name")

	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// SearchByPhoneNo in MySql
func (m *impl) SearchByPhoneNumber(ctx context.Context, phoneNumber string) ([]*contact_list_model.GlobalContactList, error) {
	var results []*contact_list_model.GlobalContactList

	// Search for the exact phone number
	if err := m.mysqlDb.
		WithContext(ctx).
		Where("phone_number = ?", phoneNumber).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
