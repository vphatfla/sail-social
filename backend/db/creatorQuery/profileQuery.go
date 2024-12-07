package creatorQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func SearchCreatorCred(id int) (model.CreatorCredential, error) {
	var creatorCred model.CreatorCredential
	err := db.DBPool.QueryRow(context.Background(), "SELECT id, email, phone_number, is_verified, is_onboarded FROM creator_credential WHERE id = $1;", id).
		Scan(&creatorCred.ID, &creatorCred.Email, &creatorCred.PhoneNumber, &creatorCred.IsVerified, &creatorCred.IsOnboarded)
	return creatorCred, err
}

func SearchCreatorInfo(id int) (model.CreatorInfo, error) {
	var creatorInfo model.CreatorInfo
	err := db.DBPool.QueryRow(context.Background(), "SELECT * FROM creator_info WHERE id = $1;", id).
		Scan(&creatorInfo.ID, &creatorInfo.Email, &creatorInfo.PhoneNumber, &creatorInfo.FirstName, &creatorInfo.LastName, &creatorInfo.Introduction, &creatorInfo.AvtLink, &creatorInfo.Address, &creatorInfo.Zipcode, &creatorInfo.City, &creatorInfo.State, &creatorInfo.Country)
	return creatorInfo, err
}
