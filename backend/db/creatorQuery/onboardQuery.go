package creatorQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewInfoRecordAndUpdateOnBoardingStatus(creatorInfo model.CreatorInfo) (int, error) {
	var id int

	tx, err := db.DBPool.Begin(context.Background())

	if err != nil {
		return id, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	err = tx.QueryRow(context.Background(), "INSERT INTO creator_info (id, email, phone_number, first_name, last_name, introduction, address, zipcode) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;",
		creatorInfo.ID, creatorInfo.Email, creatorInfo.PhoneNumber, creatorInfo.FirstName, creatorInfo.LastName, creatorInfo.Introduction, creatorInfo.Address, creatorInfo.Zipcode).Scan(&id)

	if err != nil {
		return id, err
	}

	_, err = tx.Exec(context.Background(), "UPDATE creator_credential SET is_onboarded = TRUE WHERE id = $1;", creatorInfo.ID)

	if err != nil {
		return id, err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		return id, err
	}

	return id, err
}
