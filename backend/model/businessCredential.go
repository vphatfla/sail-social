package model

type BusinessCredential struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     []byte `json:"password"`
	Phone_number string `json:"phoneNumber"`
	Is_verified  bool   `json:"isVerified"`
}
