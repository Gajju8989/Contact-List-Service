package migrate

import (
	"gorm.io/gorm"
	"task2/database/mysql/model/auth_model"
	"task2/database/mysql/model/authtoken"
	"task2/database/mysql/model/contact_list_model"
)

func MigrateAll(db *gorm.DB) error {
	if err := db.AutoMigrate(&authtoken.AuthToken{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&auth_model.UserRegistrations{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&contact_list_model.GlobalContactList{}); err != nil {
		return err
	}
	return nil
}
