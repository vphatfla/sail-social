package model

import "time"

type Post struct {
	ID         int       `json:"id,omitempty"`
	BusinessID int       `json:"businessId"`
	CreatedAt  time.Time `json:"createdAt"`
	Content    string    `json:"content"`
	PayAmount  float64   `json:"payAmount"`
	IsActive   bool      `json:"isActive"`
	WorkTime   string    `json:"workTime"`
}
