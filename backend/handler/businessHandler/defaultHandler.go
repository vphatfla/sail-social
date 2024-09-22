package businessHandler

import "net/http"

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello from default api of business!"))
}
