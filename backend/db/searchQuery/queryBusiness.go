package searchQuery

import (
	"context"
	"fmt"
	"strings"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func GetBusiness(businessType, city, state, country, zipcode *string) ([]model.BusinessInfo, error) {
	var listBusiness []model.BusinessInfo
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT * FROM business_info")

	// dynamically use the variable

	params := map[string]interface{}{
		"city":          nil,
		"state":         nil,
		"country":       nil,
		"zipcode":       nil,
		"business_type": nil,
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
	if businessType != nil {
		params["business_type"] = *businessType
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
		var businessInfo model.BusinessInfo
		if err := rows.Scan(
			&businessInfo.ID,
			&businessInfo.Email,
			&businessInfo.PhoneNumber,
			&businessInfo.FirstName,
			&businessInfo.LastName,
			&businessInfo.BusinessName,
			&businessInfo.BusinessType,
			&businessInfo.Introduction,
			&businessInfo.AvtLink,
			&businessInfo.Address,
			&businessInfo.Zipcode,
			&businessInfo.City,
			&businessInfo.State,
			&businessInfo.Country,
		); err != nil {
			return nil, err
		}
		listBusiness = append(listBusiness, businessInfo)
	}

	return listBusiness, err
}
