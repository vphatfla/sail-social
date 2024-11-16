package model

type BusinessInfo struct {
	ID           int     `json:"id"`
	Email        string  `json:"email" db:"character varying(255)"`
	PhoneNumber  string  `json:"phoneNumber" db:"character varying(20)"`
	FirstName    string  `json:"firstName" db:"character varying(100)"`
	LastName     string  `json:"lastName" db:"character varying(100)"`
	BusinessName string  `json:"businessName" db:"character varying(255)"`
	BusinessType string  `json:"businessType" db:"character varying(100)"`
	Introduction string  `json:"introduction" db:"text"`
	AvtLink      *string `json:"avtLink,omitempty" db:"text"`
	Address      string  `json:"address" db:"character varying(255)"`
	Zipcode      string  `json:"zipcode" db:"character varying(10)"`
	City         string  `json:"city" db:"character varying(255)"`
	State        string  `json:"state" db:"character varying(255)"`
	Country      string  `json:"country" db:"character varying(255)"`
}
