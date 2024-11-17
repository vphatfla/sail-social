package searchQuery

import (
	"context"
	"fmt"
	"strings"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func GetCreator(city, state, country, zipcode *string) ([]model.CreatorInfo, error) {
	var listCreator []model.CreatorInfo
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT * FROM creator_info")

	// dynamically use the variable

	params := map[string]interface{}{
		"city":    nil,
		"state":   nil,
		"country": nil,
		"zipcode": nil,
	}

	if city != nil {
		params["city"] = *city
	}
	if state != nil {
		params["state"] = *state
	}
	if country != nil {
		params["country"] = *country
	}
	if zipcode != nil {
		params["zipcode"] = *zipcode
	}

	var conditionList []string
	var paramList []any

	counter := 1
	for key, value := range params {
		if value != nil {
			conditionList = append(conditionList, fmt.Sprintf("%s=$%d ", key, counter))
			paramList = append(paramList, value)
		}
	}

	if len(conditionList) != 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditionList, " AND "))
	}

	queryBuilder.WriteString(";")

	fmt.Println(queryBuilder.String())

	rows, err := db.DBPool.Query(context.Background(), queryBuilder.String(), paramList...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var creatorInfo model.CreatorInfo
		if err := rows.Scan(
			&creatorInfo.ID,
			&creatorInfo.Email,
			&creatorInfo.PhoneNumber,
			&creatorInfo.FirstName,
			&creatorInfo.LastName,
			&creatorInfo.Introduction,
			&creatorInfo.AvtLink,
			&creatorInfo.Address,
			&creatorInfo.Zipcode,
			&creatorInfo.City,
			&creatorInfo.State,
			&creatorInfo.Country,
		); err != nil {
			return nil, err
		}
		listCreator = append(listCreator, creatorInfo)
	}

	return listCreator, err
}
