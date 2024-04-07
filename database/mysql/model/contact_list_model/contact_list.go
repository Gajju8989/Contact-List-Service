package contact_list_model

type GlobalContactList struct {
	Name        string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:20;primaryKey;"`
	Spam        string `gorm:"default:'false'"`
}
