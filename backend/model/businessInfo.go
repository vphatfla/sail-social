package model

type BusinessInfo struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	BusinessName string `json:"businessName"`
	BusinessType string `json:"businessType"`
	Introduction string `json:"introduction"`
	AvtLink      string `json:"avtLink,omitempty"`
	Address      string `json:"address"`
	Zipcode      string `json:"zipcode"`
}
