package model

type CreatorInfo struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Introduction string `json:"introduction"`
	AvtLink      string `json:"avtLink,omitempty"`
	Address      string `json:"address"`
	Zipcode      string `json:"zipcode"`
}
