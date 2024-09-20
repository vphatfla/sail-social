package model

type Post struct {
	ID          int     `json:"id"`
	Business_ID int     `json:"business_id"`
	Created_at  string  `json:"created_at"`
	Content     string  `json:"content"`
	Pay_amount  float64 `json:"pay_amount"`
	Is_active   bool    `json:"is_active"`
	Work_time   string  `json:"work_time"`
}
