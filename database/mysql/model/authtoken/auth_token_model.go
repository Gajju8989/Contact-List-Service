package authtoken

type AuthToken struct {
	PhoneNo string `gorm:"primary_key;size:20;"`
	Token   string `gorm:"size:2048;not null"`
	Expiry  int64  `gorm:"not null"`
}
