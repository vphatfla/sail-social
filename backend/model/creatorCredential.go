package model

type CreatorCredential struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	RawPassword    string `json:"password"`
	HashedPassword []byte
	PhoneNumber    string `json:"phoneNumber"`
	IsVerified     bool   `json:"isVerified"`
}
