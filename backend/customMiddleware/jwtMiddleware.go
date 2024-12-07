package customMiddleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"vphatlfa/booster-hub/auth/jwtToken"
	customError "vphatlfa/booster-hub/customError"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "No Bearer Auth Header"})
			return
		}

		// Extract the token from the header (format: "Bearer token")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "No Bearer Auth Header"})
			return
		}

		tokenString := parts[1]

		// cid := r.URL.Query().Get("cid")
		// if cid == "" {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Creator Id is empty"})
		// 	return
		// }

		var id int
		var typeU string
		id, typeU, err := jwtToken.VerifyJWTToken(tokenString)

		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error()})
			return
		}

		// Do i need a context? need to fix here
		ctx := context.WithValue(r.Context(), "default", map[string]string{
			"id":       strconv.Itoa(id),
			"typeUser": typeU,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
