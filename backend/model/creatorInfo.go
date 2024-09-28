package model

type CreatorInfo struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Phone_number string `json:"phoneNumber"`
	First_name   string `json:"firstName"`
	Last_name    string `json:"lastName"`
	Introduction string `json:"introduction"`
	Avt_link     string `json:"avtLink,omitempty"`
	Address      string `json:"address"`
	Zipcode      string `json:"zipcode"`
}
