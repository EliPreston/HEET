package services

import (
	"context"
	"fmt"
	"os"

	// For connecting to Postgres DB
	"github.com/jackc/pgx/v5"
)

var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbUser = os.Getenv("DB_USER")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbName = os.Getenv("DB_NAME")
var connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

func ConnectToDB() {

	fmt.Println(connStr)

	var conn, err = pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var init string
	err = conn.QueryRow(context.Background(), "select 'Connected to database...'").Scan(&init)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(2)
	}

	fmt.Println(init)
}
