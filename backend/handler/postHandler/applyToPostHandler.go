package postHandler

import (
	"encoding/json"
	"net/http"
	customeError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/postQuery"
	"vphatlfa/booster-hub/model"
)

func NewApplyToPostHandler(w http.ResponseWriter, r *http.Request) {
	var creatorPostApplied model.CreatorPostApplied
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&creatorPostApplied)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	creatorId, postId, err := postQuery.InsertNewApplyRecord(creatorPostApplied)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Data base faield : " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"creatorId": creatorId, "postId": postId})
}

func UpdateApplyToPostHandler(w http.ResponseWriter, r *http.Request) {
	var creatorPostApplied model.CreatorPostApplied

	err := json.NewDecoder(r.Body).Decode(&creatorPostApplied)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	creatorId, postId, err := postQuery.UpdateApplyRecord(creatorPostApplied)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Data base faield : " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"creatorId": creatorId, "postId": postId})
}
