package searchQuery

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func GetAllPost() ([]model.Post, error) {
	var listPost []model.Post

	rows, err := db.DBPool.Query(context.Background(), "SELECT id, business_id, created_at, content, pay_amount, is_active, work_time FROM post;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.ID,
			&post.BusinessID,
			&post.CreatedAt,
			&post.Content,
			&post.PayAmount,
			&post.IsActive,
			&post.WorkTime,
		); err != nil {
			return nil, err
		}
		listPost = append(listPost, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPost, err
}

func GetAllPostWithCondition(bid *int, city, state, country, zipcode *string) ([]model.Post, error) {
	var listPost []model.Post
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT p.* FROM post p JOIN business_info b ON p.business_id = b.id")

	// dynamically use the variable

	params := map[string]interface{}{
		"bid":     bid,
		"city":    city,
		"state":   state,
		"country": country,
		"zipcode": zipcode,
	}

	var conditionList []string
	var paramList []any

	fmt.Println(params["bid"] == nil)
	counter := 0
	for key, value := range params {
		if !reflect.ValueOf(value).IsNil() {
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

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.ID,
			&post.BusinessID,
			&post.CreatedAt,
			&post.Content,
			&post.PayAmount,
			&post.IsActive,
			&post.WorkTime,
		); err != nil {
			return nil, err
		}
		listPost = append(listPost, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPost, err
}
