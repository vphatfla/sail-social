package businessQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewInfoRecord(businessInfo model.BusinessInfo) (int, error) {
	var id int
	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO business_info (id, email, phone_number, first_name, last_name, introduction, address, zipcode, business_name, business_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;",
		businessInfo.ID, businessInfo.Email, businessInfo.PhoneNumber, businessInfo.FirstName, businessInfo.LastName, businessInfo.Introduction, businessInfo.Address, businessInfo.Zipcode, businessInfo.BusinessName, businessInfo.BusinessType).Scan(&id)
	return id, err
}
