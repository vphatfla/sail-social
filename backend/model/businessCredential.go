package model

type BusinessCredential struct {
	ID             int64  `json:"id,omitempty"`
	Email          string `json:"email"`
	RawPassword    string `json:"password"`
	HashedPassword []byte
	PhoneNumber    string `json:"phoneNumber,omitempty"`
	IsVerified     bool   `json:"isVerified,omitempty"`
}
