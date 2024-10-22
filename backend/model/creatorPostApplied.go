package model

type CreatorPostApplied struct {
	CreatorId int    `json:"creatorId"`
	PostId    int    `json:"postId"`
	Message   string `json:"message"`
}
