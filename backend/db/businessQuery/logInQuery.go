package businessQuery

import (
	"context"
	"database/sql"
	"fmt"
	"vphatlfa/booster-hub/db"
)

func GetHashPassword(email string) ([]byte, error) {
	var hashedPassword []byte
	err := db.DBPool.QueryRow(context.Background(), "SELECT password FROM business_credential WHERE email =$1", email).Scan(&hashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return hashedPassword, fmt.Errorf("email does not exist in the db")
		}
		return hashedPassword, fmt.Errorf("error %s", err.Error())
	}

	return hashedPassword, nil
}
