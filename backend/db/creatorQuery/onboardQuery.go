package creatorQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewInfoRecord(creatorInfo model.CreatorInfo) (int, error) {
	var id int
	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO creator_info (id, email, phone_number, first_name, last_name, introduction, address, zipcode) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;",
		creatorInfo.ID, creatorInfo.Email, creatorInfo.PhoneNumber, creatorInfo.FirstName, creatorInfo.LastName, creatorInfo.Introduction, creatorInfo.Address, creatorInfo.Zipcode).Scan(&id)
	return id, err
}
