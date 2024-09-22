package model

type BusinessInfo struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Phone_number   string `json:"phoneNumber"`
	First_name     string `json:"firstName"`
	Last_name      string `json:"lastName"`
	Business_name  string `json:"businessName"`
	Business_type  string `json:"businessType"`
	Introduction   string `json:"introduction"`
	Avt_link       string `json:"avtLink,omitempty"`
	Street_address string `json:"streetAddress"`
	Zipcode        string `json:"zipcode"`
}
