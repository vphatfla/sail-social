package businessHandler

import (
	"encoding/json"
	"net/http"
	"vphatlfa/booster-hub/auth/hashing"
	"vphatlfa/booster-hub/auth/jwtToken"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/businessQuery"
	"vphatlfa/booster-hub/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var businessCredential model.BusinessCredential
	// Decode the json payload
	err := json.NewDecoder(r.Body).Decode(&businessCredential)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	hashedPassword, err := businessQuery.GetHashPassword(businessCredential.Email)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error()})
		return
	}

	err = hashing.VerifyPassword(hashedPassword, businessCredential.RawPassword)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " passwords do not match"})
		return
	}

	businessCredential, err = businessQuery.GetBusinessCredentialByEmail(businessCredential.Email)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " passwords do not match"})
		return
	}

	// token
	token, err := jwtToken.GenerateJWTToken(businessCredential.ID, "business")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Token : " + err.Error()})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
