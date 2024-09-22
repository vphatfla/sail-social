package model

type CreatorPortfolio struct {
	ID                 int    `json:"id"`
	Instagram_link     string `json:"instagramLink"`
	Tiktok_link        string `json:"tiktokLink"`
	Instagram_follower int    `json:"instagramFollower"`
	Tiktok_follower    int    `json:"tiktokFollower"`
}
