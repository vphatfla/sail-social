package creatorHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/creatorQuery"
)

type temp struct {
	ID int `json:"userId"`
}

func SearchCreatorCredInfoHandler(w http.ResponseWriter, r *http.Request) {
	var rBody temp
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&rBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	creatorCred, err := creatorQuery.SearchCreatorCred(rBody.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: " error searching " + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creatorCred)
}

func SearchCreatorInfoHandler(w http.ResponseWriter, r *http.Request) {
	var rBody temp
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&rBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	creatorInfo, err := creatorQuery.SearchCreatorInfo(rBody.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: fmt.Sprintf("error query database: %w", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creatorInfo)
}
