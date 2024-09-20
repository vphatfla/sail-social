package handler

import "net/http"

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World! Testing Air"))
}
