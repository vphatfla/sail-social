package model

import "time"

type Post struct {
	ID         int       `json:"id,omitempty" db:"id"`
	BusinessID int       `json:"businessId" db:"business_id"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	Content    string    `json:"content" db:"content"`
	PayAmount  float64   `json:"payAmount" db:"pay_amount"`
	IsActive   bool      `json:"isActive" db:"is_active"`
	WorkTime   string    `json:"workTime" db:"work_time"`
}
