package creatorHandler

import (
	"encoding/json"
	"net/http"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/creatorQuery"
	"vphatlfa/booster-hub/model"
)

func PostNewPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	var creatorPortfolio model.CreatorPortfolio
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&creatorPortfolio)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	id, err := creatorQuery.InsertNewPortfolioRecord(creatorPortfolio)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: " insert record info " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
