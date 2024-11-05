package businessQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewInfoRecordAndUpdateOnBoardingStatus(businessInfo model.BusinessInfo) (int, error) {
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

	err = tx.QueryRow(context.Background(), "INSERT INTO business_info (id, email, phone_number, first_name, last_name, introduction, address, zipcode, business_name, business_type, city, state, country) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id;",
		businessInfo.ID, businessInfo.Email, businessInfo.PhoneNumber, businessInfo.FirstName, businessInfo.LastName, businessInfo.Introduction, businessInfo.Address, businessInfo.Zipcode, businessInfo.BusinessName, businessInfo.BusinessType, businessInfo.City, businessInfo.State, businessInfo.Country).Scan(&id)

	if err != nil {
		return id, err
	}

	_, err = tx.Exec(context.Background(), "UPDATE business_credential SET is_onboarded = TRUE WHERE id = $1;", businessInfo.ID)

	if err != nil {
		return id, err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		return id, err
	}

	return id, err
}
