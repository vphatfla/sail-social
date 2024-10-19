package businessQuery

import (
	"context"
	"database/sql"
	"fmt"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func GetBusinessCredentialByEmail(email string) (model.BusinessCredential, error) {
	var businessCredential model.BusinessCredential
	err := db.DBPool.QueryRow(context.Background(), "SELECT id, email, password, phone_number, is_verified FROM business_credential WHERE email = $1;", email).Scan(&businessCredential.ID, &businessCredential.Email, &businessCredential.HashedPassword, &businessCredential.PhoneNumber, &businessCredential.IsVerified)
	if err != nil {
		if err != sql.ErrNoRows {
			return businessCredential, fmt.Errorf("error - no record found by email")
		}
		return businessCredential, err
	}
	return businessCredential, nil
}

func IsInfoRecordValid(id int, email string, phoneNumber string) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM business_credential WHERE id = $1 AND email = $2 AND phone_number = $3);", id, email, phoneNumber).Scan(&check)
	return check, err
}
