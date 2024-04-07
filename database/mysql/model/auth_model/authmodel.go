package auth_model

type UserRegistrations struct {
	Name        string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:20;primaryKey;not null"`
	Email       string `gorm:"size:100"`
	Password    string `gorm:"size:100;not null"`
}
