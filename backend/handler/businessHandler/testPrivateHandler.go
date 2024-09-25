package businessHandler

import (
	"encoding/json"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"Message": "welcome to private route!"})
}
