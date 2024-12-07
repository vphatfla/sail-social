package postHandler

import (
	"encoding/json"
	"net/http"
	customeError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/postQuery"
	"vphatlfa/booster-hub/model"
)

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	defer r.Body.Close()
	// check unvalid id, request json must include valid post id

	post.ID = -1
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return
	}

	if post.ID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Json payload invalid, no post id"})
		return
	}

	id, err := postQuery.UpdatePostQuery(post)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "Database Insert Opertion Failed" + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
