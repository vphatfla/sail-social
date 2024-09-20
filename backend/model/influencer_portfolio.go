package model

type Influencer_portfolio struct {
	ID                 int    `json:"id"`
	Instagram_link     string `json:"instagram_link"`
	Tiktok_link        string `json:"tiktok_link"`
	Instagram_follower int    `json:"instagram_follower"`
	Tiktok_follower    int    `json:"tiktok_follower"`
}
