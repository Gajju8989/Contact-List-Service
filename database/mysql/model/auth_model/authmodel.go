package auth_model

type UserRegistrations struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:20;uniqueIndex"`
	Email       string `gorm:"size:100"`
	Password    string `gorm:"size:100;not null"`
}
