package model

type BusinessCredential struct {
	ID             int    `json:"id,omitempty" db:"id"`
	Email          string `json:"email" db:"email"`
	RawPassword    string `json:"password"`
	HashedPassword []byte `db:"password"`
	PhoneNumber    string `json:"phoneNumber,omitempty" db:"phone_number"`
	IsVerified     bool   `json:"isVerified,omitempty" db:"is_verified"`
	IsOnboarded    bool   `db:"is_onboarded"`
}
