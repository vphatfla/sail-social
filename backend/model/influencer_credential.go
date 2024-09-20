package model

type Influencer_credential struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     []byte `json:"password"`
	Phone_number string `json:"phone_number"`
	Is_verified  bool   `json:"is_verified"`
}
