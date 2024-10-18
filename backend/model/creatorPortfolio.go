package model

type CreatorPortfolio struct {
	ID                int    `json:"id"`
	InstagramLink     string `json:"instagramLink,omitempty"`
	TiktokLink        string `json:"tiktokLink,omitempty"`
	InstagramFollower int    `json:"instagramFollower,omitempty"`
	TiktokFollower    int    `json:"tiktokFollower,omitempty"`
}
