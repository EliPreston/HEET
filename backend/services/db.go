package services

import (
	"context"
	"fmt"
	"os"

	// For connecting to Postgres DB
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetDBURL() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return dbStr
}

var Ctx = context.Background()
var Pool, Err_global = pgxpool.New(Ctx, SetDBURL())

func ConnectToDB() {

	if Err_global != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", Err_global)
		os.Exit(1)
	}

	if Err_global := Pool.Ping(Ctx); Err_global != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", Err_global)
		os.Exit(2)
	}

	fmt.Println("Connected to database...")
}
