package businessHandler

import (
	"encoding/json"
	"net/http"
	"vphatlfa/booster-hub/auth/jwtToken"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/businessQuery"
	"vphatlfa/booster-hub/model"
)

func OnboadingHandler(w http.ResponseWriter, r *http.Request) {
	var businessInfo model.BusinessInfo
	// decode body request json payload
	err := json.NewDecoder(r.Body).Decode(&businessInfo)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	// check if the info is legit

	check, err := businessQuery.IsInfoRecordValid(businessInfo.ID, businessInfo.Email, businessInfo.PhoneNumber)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	if !check {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Info not match"})
		return
	}

	// insert record
	// update is_onboarded record

	id, err := businessQuery.InsertNewInfoRecordAndUpdateOnBoardingStatus(businessInfo)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: " insert record info " + err.Error()})
		return
	}

	// token
	token, err := jwtToken.GenerateJWTToken(id, "business", true)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Token : " + err.Error()})
		return
	}

	response := map[string]interface{}{
		"id":    id,
		"token": token,
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
