package creatorQuery

import (
	"context"
	"database/sql"
	"fmt"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func GetCreatorCredentialByEmail(email string) (model.CreatorCredential, error) {
	var creatorCredential model.CreatorCredential
	err := db.DBPool.QueryRow(context.Background(), "SELECT id, email, password, phone_number, is_verified FROM creator_credential WHERE email = $1;", email).Scan(&creatorCredential.ID, &creatorCredential.Email, &creatorCredential.HashedPassword, &creatorCredential.PhoneNumber, &creatorCredential.IsVerified)
	if err != nil {
		if err != sql.ErrNoRows {
			return creatorCredential, fmt.Errorf("error - no record found by email")
		}
		return creatorCredential, err
	}
	return creatorCredential, nil
}

func IsInfoRecordValid(id int, email string, phoneNumber string) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM creator_credential WHERE id = $1 AND email = $2 AND phone_number = $3);", id, email, phoneNumber).Scan(&check)
	return check, err
}
