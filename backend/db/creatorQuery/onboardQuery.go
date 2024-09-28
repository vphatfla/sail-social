package creatorQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewInfoRecord(creatorInfo model.CreatorInfo) (int64, error) {
	var id int64
	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO creator_info (id, email, phone_number, first_name, last_name, introduction, address, zipcode) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;",
		creatorInfo.ID, creatorInfo.Email, creatorInfo.Phone_number, creatorInfo.First_name, creatorInfo.Last_name, creatorInfo.Introduction, creatorInfo.Address, creatorInfo.Zipcode).Scan(&id)
	return id, err
}

func IsInfoRecordValid(creatorInfo model.CreatorInfo) (bool, error) {
	var check bool
	err := db.DBPool.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM creator_credential WHERE id = $1 AND email = $2 AND phone_number = $3);", creatorInfo.ID, creatorInfo.Email, creatorInfo.Phone_number).Scan(&check)
	return check, err
}
