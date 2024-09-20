package main

import (
	"fmt"
	"vphatlfa/booster-hub/db"
)

func main() {
	fmt.Print("App Started\n\n")

	db.InitPostgresPoolConnection()

}
