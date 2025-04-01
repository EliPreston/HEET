package services

import (
	"context"
	"fmt"
	"os"

	// For connecting to Postgres DB
	"github.com/jackc/pgx/v5"
)

// docker run --rm --name docker_container_img_name --env POSTGRES_PASSWORD=password --env POSTGRES_DB=heet --volume pg-data:/var/lib/postgresql/data --publish 5432:5432 postgres:bookworm
// docker exec -it docker_container_img_name psql -U postgres

var dbURL = "postgres://postgres:admin@localhost:5432/gotodo"
var conn, err = pgx.Connect(context.Background(), dbURL)

func ConnectToDB() {

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
