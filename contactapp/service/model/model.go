package model

type UserRegistrationsRequestModel struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"-"`
}

type ContactListRequest struct {
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Requester   Request `json:"requester"`
}
type Request struct {
	RequesterPhoneNo string `json:"requester_phone_no"`
}
