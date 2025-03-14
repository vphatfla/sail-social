package searchQuery

import (
	"context"
	"fmt"
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
		"business_id": nil,
		"city":        nil,
		"state":       nil,
		"country":     nil,
		"zipcode":     nil,
	}

	if bid != nil {
		params["business_id"] = *bid
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

	return listPost, err
}

func GetPostAppliedByCreator(cid int) ([]model.CreaterPostApplied, error) {
	var listPost []model.CreaterPostApplied
	query := "SELECT p.*, cpa.message FROM post p JOIN creator_post_applied cpa ON p.id = cpa.post_id WHERE cpa.creator_id = $1;"

	rows, err := db.DBPool.Query(context.Background(), query, cid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var post model.CreaterPostApplied
		if err := rows.Scan(
			&post.ID,
			&post.BusinessID,
			&post.CreatedAt,
			&post.Content,
			&post.PayAmount,
			&post.IsActive,
			&post.WorkTime,
			&post.Message,
		); err != nil {
			return nil, err
		}
		listPost = append(listPost, post)
	}

	return listPost, nil
}

func GetPostSavedByCreator(cid int) ([]model.Post, error) {
	var listPost []model.Post
	query := "SELECT * FROM post p JOIN creator_post_saved cps ON cps.post_id = p.id WHERE cps.creator_id = $1;"

	rows, err := db.DBPool.Query(context.Background(), query, cid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

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

	return listPost, nil
}

func GetPostByBusiness(bid int) ([]model.Post, error) {
	var listPost []model.Post
	query := "SELECT * FROM post WHERE business_id = $1;"

	rows, err := db.DBPool.Query(context.Background(), query, bid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

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

	return listPost, nil
}
