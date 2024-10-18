package model

type Post struct {
	ID         int     `json:"id"`
	BusinessID int     `json:"businessId"`
	CreatedAt  string  `json:"createdAt"`
	Content    string  `json:"content"`
	PayAmount  float64 `json:"payAmount"`
	IsActive   bool    `json:"isActive"`
	WorkTime   string  `json:"workTime"`
}
