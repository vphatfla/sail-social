package creatorHandler

import (
	"encoding/json"
	"net/http"
	"vphatlfa/booster-hub/auth/hashing"
	customeError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/creatorQuery"
	"vphatlfa/booster-hub/model"
)

func CreatorSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var creatorCredential model.CreatorCredential

	// Decode the json payload
	err := json.NewDecoder(r.Body).Decode(&creatorCredential)

	if err != nil {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	// check email exist
	check, err := creatorQuery.CheckIfEmailExists(creatorCredential.Email)

	if err != nil {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " error checking email not OK"})
		return
	}

	if check {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Email Exists"})
		return
	}

	// check phone exist
	check, err = creatorQuery.CheckIfPhoneNumberExists(creatorCredential.PhoneNumber)

	if err != nil {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " error checking phone not OK"})
		return
	}

	if check {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Phone Number Exists"})
		return
	}

	// hash operation
	hash, err := hashing.HashPassword(string(creatorCredential.RawPassword))
	if err != nil {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Hashing not OK" + err.Error()})
		return
	}
	creatorCredential.HashedPassword = hash
	// insert the record
	err = creatorQuery.InsertNewRecord(creatorCredential)
	if err != nil {
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Database : " + err.Error()})
		return
	}
}
