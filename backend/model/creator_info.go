package model

type Creator_info struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Introduction string `json:"introduction"`
	Avt_link     string `json:"avt_link,omitempty"`
	Address      string `json:"address"`
	Zipcode      string `json:"zipcode"`
}
