package searchQuery

import (
	"context"
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

	query := `
		SELECT p.*
		FROM post p
		JOIN business_info b ON p.business_id = b.id
		WHERE (COALESCE($1, p.business_id) = p.business_id)
		AND (COALESCE($2, b.city) = b.city)
		AND (COALESCE($3, b.state) = b.state)
		AND (COALESCE($4, b.county) = b.country)
		AND (COALESCE($5, b.zipcode) = b.zipcode)

	`
	rows, err := db.DBPool.Query(context.Background(), query, bid, city, state, country, zipcode)
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
