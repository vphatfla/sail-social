package model

type Business_info struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Phone_number   string `json:"phone_number"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Business_name  string `json:"business_name"`
	Business_type  string `json:"business_type"`
	Introduction   string `json:"introduction"`
	Avt_link       string `json:"avt_link,omitempty"`
	Street_address string `json:"street_address"`
	Zipcode        string `json:"zipcode"`
}
