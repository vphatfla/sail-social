package businessQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func SearchBusinessCred(id int) (model.BusinessCredential, error) {
	var businessCred model.BusinessCredential
	err := db.DBPool.QueryRow(context.Background(), "SELECT id, email, phone_number, is_verified, is_onboarded FROM business_credential WHERE id = $1;", id).
		Scan(&businessCred.ID, &businessCred.Email, &businessCred.PhoneNumber, &businessCred.IsVerified, &businessCred.IsOnboarded)
	return businessCred, err
}

func SearchBusinessInfo(id int) (model.BusinessInfo, error) {
	var businessInfo model.BusinessInfo
	err := db.DBPool.QueryRow(context.Background(), "SELECT * FROM business_info WHERE id = $1;", id).
		Scan(&businessInfo.ID, &businessInfo.Email, &businessInfo.PhoneNumber, &businessInfo.FirstName, &businessInfo.LastName, &businessInfo.BusinessName, &businessInfo.BusinessType, &businessInfo.Introduction, &businessInfo.AvtLink, &businessInfo.Address, &businessInfo.Zipcode, &businessInfo.City, &businessInfo.State, &businessInfo.Country)
	return businessInfo, err
}
