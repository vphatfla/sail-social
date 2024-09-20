package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DBPool *pgxpool.Pool
)

func InitPostgresPoolConnection() {
	fmt.Println("Creating DB Connection Pool")
	var err error
	DBPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool to Postgresql: %v\n", err)
		os.Exit(1)
	}
	err = DBPool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping connection pool to Postgresql: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Connect to Postgresql successfully!\n\n")
	// Defer to close the db pool later
	defer DBPool.Close()
}
