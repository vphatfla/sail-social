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

func businessSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var businessCredential model.BusinessCredential

	// Decode the json payload
	err := json.NewDecoder(r.Body).Decode(&businessCredential)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	// check email exist
	check, err := businessQuery.CheckIfEmailExists(businessCredential.Email)

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
	check, err = businessQuery.CheckIfPhoneNumberExists(businessCredential.PhoneNumber)

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
	hash, err := hashing.HashPassword(string(businessCredential.RawPassword))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Hashing not OK" + err.Error()})
		return
	}
	businessCredential.HashedPassword = hash
	// insert the record
	id, err := businessQuery.InsertNewRecord(businessCredential)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Database : " + err.Error()})
		return
	}

	token, err := jwtToken.GenerateJWTToken(id, "business")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Token : " + err.Error()})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
