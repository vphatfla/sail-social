package customMiddleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"vphatlfa/booster-hub/auth/jwtToken"
	customError "vphatlfa/booster-hub/customError"
)

func CreatorJWTMiddleware(next http.Handler) http.Handler {
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

		cid := r.URL.Query().Get("cid")
		if cid == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "Creator Id is empty"})
			return
		}

		check, err := jwtToken.VerifyJWTToken(tokenString, cid, "creator")

		if err != nil || !check {
			w.WriteHeader(401)
			if err != nil {
				json.NewEncoder(w).Encode(customError.ErrorMessage{Message: err.Error()})
			} else {
				json.NewEncoder(w).Encode(customError.ErrorMessage{Message: "JWT Check failed"})
			}

			return
		}

		// Do i need a context?
		// ctx := context.WithValue(r.Context(), "default", map[string]string{
		// 	"cid":      cid,
		// 	"typeUser": "creator",
		// })

		next.ServeHTTP(w, r)
	})
}
