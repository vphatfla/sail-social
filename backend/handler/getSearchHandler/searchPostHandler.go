package getSearchHandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	customeError "vphatlfa/booster-hub/customError"
	"vphatlfa/booster-hub/db/searchQuery"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil && !errors.Is(err, io.EOF) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error() + " json payload invalid"})
		return

	}
	defer r.Body.Close()

	var bid *int
	var city, state, country, zipcode *string

	if x, ok := payload["businessId"].(float64); ok {
		value := int(x)
		bid = &value
	}

	if x, ok := payload["city"].(string); ok {
		city = &x
	}

	if x, ok := payload["state"].(string); ok {
		state = &x
	}

	if x, ok := payload["country"].(string); ok {
		country = &x
	}

	if x, ok := payload["zipcode"].(string); ok {
		zipcode = &x
	}

	listPost, err := searchQuery.GetAllPostWithCondition(bid, city, state, country, zipcode)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(listPost)
}

func GetPostAppliedByCreatorHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}
	var cid int

	json.NewDecoder(r.Body).Decode(&payload)

	if payload["creatorId"] == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "json payload invalid"})
		return
	}
	if f, ok := payload["creatorId"].(float64); ok {
		cid = int(f)
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "creatorId is not valid"})
		return
	}

	listPost, err := searchQuery.GetPostAppliedByCreator(cid)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(listPost)
}

func GetPostSavedByCreatorHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}
	var cid int

	json.NewDecoder(r.Body).Decode(&payload)

	if payload["creatorId"] == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "json payload invalid"})
		return
	}

	if f, ok := payload["creatorId"].(float64); ok {
		cid = int(f)
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "creatorId is not valid"})
		return
	}

	listPost, err := searchQuery.GetPostSavedByCreator(cid)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(listPost)
}

func GetPostByBusinessHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}
	var bid int

	json.NewDecoder(r.Body).Decode(&payload)

	if payload["businessId"] == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "json payload invalid"})
		return
	}

	if f, ok := payload["businessId"].(float64); ok {
		bid = int(f)
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: "creatorId is not valid"})
		return
	}

	listPost, err := searchQuery.GetPostByBusiness(bid)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(customeError.ErrorMessage{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(listPost)
}
