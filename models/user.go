package models

type User struct {
	Id          int64  `json:"id"`
	DeviceId    string `json:"device_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email_name"`
	CreatedDate string `json:"created_date"`
	Sex         string `json:"sex"`
}
