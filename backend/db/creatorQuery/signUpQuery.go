package creatorQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func CheckIfEmailExists(email string) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM creator_credential WHERE email = $1);", email).Scan(&check)
	return check, err
}

func CheckIfPhoneNumberExists(phoneNumber string) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM creator_credential WHERE phone_number = $1);", phoneNumber).Scan(&check)
	return check, err
}

func InsertNewRecord(creator model.CreatorCredential) error {
	_, err := db.DBPool.Exec(context.Background(), "INSERT INTO creator_credential (email, password, phone_number, is_verified) VALUES ($1, $2, $3, $4);",
		creator.Email, creator.HashedPassword, creator.PhoneNumber, false)
	return err
}
