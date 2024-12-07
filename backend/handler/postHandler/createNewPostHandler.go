package postHandler

import (
	"encoding/json"
	"net/http"
	customeError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/postQuery"
	"vphatlfa/booster-hub/model"
)

func CreateNewPostHandler(w http.ResponseWriter, r *http.Request) {
	var newPost model.Post
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&newPost)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	id, err := postQuery.InsertNewPostRecord(newPost)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Database Insert Opertion Failed" + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
