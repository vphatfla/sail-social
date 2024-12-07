package businessHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/businessQuery"
)

type temp struct {
	ID int `json:"userId"`
}

func SearchBusinessCredHandler(w http.ResponseWriter, r *http.Request) {
	var rBody temp
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&rBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	businessCred, err := businessQuery.SearchBusinessCred(rBody.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: " error searching " + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(businessCred)
}

func SearchBusinessInfoHandler(w http.ResponseWriter, r *http.Request) {
	var rBody temp
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&rBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	businessInfo, err := businessQuery.SearchBusinessInfo(rBody.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: fmt.Sprintf("error query database: %w", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(businessInfo)
}
