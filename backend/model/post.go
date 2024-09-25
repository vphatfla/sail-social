package model

type Post struct {
	ID          int     `json:"id"`
	Business_ID int     `json:"businessId"`
	Created_at  string  `json:"createdAt"`
	Content     string  `json:"content"`
	Pay_amount  float64 `json:"payAmount"`
	Is_active   bool    `json:"isActive"`
	Work_time   string  `json:"workTime"`
}
