package main

import (
	"fmt"
	"net/http"
	"vphatlfa/booster-hub/db"
	"vphatlfa/booster-hub/router"
)

func main() {
	fmt.Print("App Started\n\n")

	db.InitPostgresPoolConnection()
	// Defer to close the db pool later
	defer db.DBPool.Close()

	http.ListenAndServe(":3000", router.MainRouter())
}
