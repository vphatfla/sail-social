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

func CreatorSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var creatorCredential model.CreatorCredential

	// Decode the json payload
	err := json.NewDecoder(r.Body).Decode(&creatorCredential)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	// check email exist
	check, err := creatorQuery.CheckIfEmailExists(creatorCredential.Email)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " error checking email not OK"})
		return
	}

	if check {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Email Exists"})
		return
	}

	// check phone exist
	check, err = creatorQuery.CheckIfPhoneNumberExists(creatorCredential.PhoneNumber)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " error checking phone not OK"})
		return
	}

	if check {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Phone Number Exists"})
		return
	}

	// hash operation
	hash, err := hashing.HashPassword(string(creatorCredential.RawPassword))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Hashing not OK" + err.Error()})
		return
	}
	creatorCredential.HashedPassword = hash
	// insert the record
	id, err := creatorQuery.InsertNewCredentialRecord(creatorCredential)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Database : " + err.Error()})
		return
	}

	token, err := jwtToken.GenerateJWTToken(id, "creator")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Token : " + err.Error()})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
