package postQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewApplyRecord(creatorPostApplied model.CreatorPostApplied) (int, int, error) {
	var creatorId int
	var postId int

	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO creator_post_applied (creator_id, post_id, message) VALUES ($1, $2, $3) RETURNING creator_id, post_id;",
		creatorPostApplied.CreatorId, creatorPostApplied.PostId, creatorPostApplied.Message).Scan(&creatorId, &postId)

	return creatorId, postId, err
}

func UpdateApplyRecord(creatorPostApplied model.CreatorPostApplied) (int, int, error) {
	var creatorId int
	var postId int

	err := db.DBPool.QueryRow(context.Background(), "UPDATE creator_post_applied SET message = $1 WHERE creator_id = $2 AND post_id = $3 RETURNING creator_id, post_id;",
		creatorPostApplied.Message, creatorPostApplied.CreatorId, creatorPostApplied.PostId).Scan(&creatorId, &postId)

	return creatorId, postId, err
}
