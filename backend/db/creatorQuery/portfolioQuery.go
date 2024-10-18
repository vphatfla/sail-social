package creatorQuery

import (
	"context"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/model"
)

func InsertNewPortfolioRecord(creatorPortfolio model.CreatorPortfolio) (int, error) {
	var id int
	err := db.DBPool.QueryRow(context.Background(), "INSERT INTO creator_portfolio (id, instagram_link, tiktok_link) VALUES ($1, $2, $3) RETURNING id;", creatorPortfolio.ID, creatorPortfolio.InstagramLink, creatorPortfolio.TiktokLink).Scan(&id)
	return id, err
}
