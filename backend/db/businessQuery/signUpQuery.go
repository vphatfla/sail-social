package businessQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func CheckIfEmailExists(email string) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM business_credential WHERE email = $1);", email).Scan(&check)
	return check, err
}

func CheckIfPhoneNumberExists(phoneNumber string) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM business_credential WHERE phone_number = $1);", phoneNumber).Scan(&check)
	return check, err
}

func InsertNewRecord(business model.BusinessCredential) (int, error) {
	var id int
	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO business_credential (email, password, phone_number, is_verified) VALUES ($1, $2, $3, $4) RETURNING id;",
		business.Email, business.HashedPassword, business.PhoneNumber, false).Scan(&id)
	return id, err
}
