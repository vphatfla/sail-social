package postQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewPostRecord(newPost model.Post) (int, error) {
	var id int
	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO post (business_id, created_at, content, pay_amount, is_active, work_time) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;",
		newPost.BusinessID, newPost.CreatedAt, newPost.Content, newPost.PayAmount, newPost.IsActive, newPost.WorkTime).Scan(&id)
	return id, err
}

func UpdatePostQuery(post model.Post) (int, error) {
	var id int
	err := db.DBPool.QueryRow(context.Background(), "UPDATE post SET content = $1, pay_amount = $2, is_active = $3, work_time= $4 WHERE id = $5 RETURNING id;",
		post.Content, post.PayAmount, post.IsActive, post.WorkTime, post.ID).Scan(&id)
	return id, err
}
