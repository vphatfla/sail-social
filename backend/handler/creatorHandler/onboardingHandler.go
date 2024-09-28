package creatorHandler

import (
	"encoding/json"
	"net/http"
	customError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/model"
)

func OnboadingHandler(w http.ResponseWriter, r *http.Request) {
	var creatorInfo model.CreatorInfo
	// decode body request json paylaod
	err := json.NewDecoder(r.Body).Decode(&creatorInfo)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}
}
