package creatorHandler

import (
	"encoding/json"
	"net/http"
	"vphatlfa/booster-hub/auth/jwtToken"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/creatorQuery"
	"vphatlfa/booster-hub/model"
)

func OnboadingHandler(w http.ResponseWriter, r *http.Request) {
	var creatorInfo model.CreatorInfo
	// decode body request json payload
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&creatorInfo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	// check if the info is legit

	check, err := creatorQuery.IsInfoRecordValid(int(creatorInfo.ID), creatorInfo.Email, creatorInfo.PhoneNumber)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	if !check {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Info not match"})
		return
	}

	// insert record

	id, err := creatorQuery.InsertNewInfoRecordAndUpdateOnBoardingStatus(creatorInfo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: " insert record info " + err.Error()})
		return
	}

	// token
	token, err := jwtToken.GenerateJWTToken(id, "creator", true)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Token : " + err.Error()})
		return
	}

	response := map[string]interface{}{
		"id":    id,
		"token": token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
