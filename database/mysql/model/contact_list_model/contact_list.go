package contact_list_model

type ContactList struct {
	ID          uint   `gorm:"primaryKey"` //this will be the foreign key for UserRegistrations table
	Name        string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:20;not null"`
	Spam        bool   `gorm:"default:false"`
}
