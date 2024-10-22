package creatorHandler

import (
	"encoding/json"
	"net/http"
	"vphatlfa/booster-hub/auth/hashing"
	"vphatlfa/booster-hub/auth/jwtToken"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/creatorQuery"
	"vphatlfa/booster-hub/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creatorCredential model.CreatorCredential
	// Decode the json payload
	err := json.NewDecoder(r.Body).Decode(&creatorCredential)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	hashedPassword, err := creatorQuery.GetHashPassword(creatorCredential.Email)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error()})
		return
	}

	err = hashing.VerifyPassword(hashedPassword, creatorCredential.RawPassword)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " passwords do not match"})
		return
	}

	creatorCredential, err = creatorQuery.GetCreatorCredentialByEmail(creatorCredential.Email)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " passwords do not match"})
		return
	}

	// token
	token, err := jwtToken.GenerateJWTToken(creatorCredential.ID, "creator", creatorCredential.IsOnboarded)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Token : " + err.Error()})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
