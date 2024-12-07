package model

type CreatorCredential struct {
	ID             int    `json:"id,omitempty" db:"id"`
	Email          string `json:"email" db:"email"`
	RawPassword    string `json:"password,omitempty"`
	HashedPassword []byte `json:"-" db:"password"`
	PhoneNumber    string `json:"phoneNumber" db:"phone_number"`
	IsVerified     bool   `json:"isVerified" db:"is_verified"`
	IsOnboarded    bool   `db:"is_onboarded"`
}
