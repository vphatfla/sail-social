package creatorHandler

import (
	"encoding/json"
	"net/http"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/creatorQuery"
	"vphatlfa/booster-hub/model"
)

func OnboadingHandler(w http.ResponseWriter, r *http.Request) {
	var creatorInfo model.CreatorInfo
	// decode body request json payload
	err := json.NewDecoder(r.Body).Decode(&creatorInfo)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	// check if the info is legit

	check, err := creatorQuery.IsInfoRecordValid(creatorInfo)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	if check {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Info not match"})
		return
	}

	// insert record

	id, err := creatorQuery.InsertNewInfoRecord(creatorInfo)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: " insert record info " + err.Error()})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}
