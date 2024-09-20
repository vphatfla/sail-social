package main

import (
	"fmt"
	"net/http"
	"vphatlfa/booster-hub/db"
)

func main() {
	fmt.Print("App Started\n\n")

	db.InitPostgresPoolConnection()

	http.ListenAndServe(":3000", router())
}
