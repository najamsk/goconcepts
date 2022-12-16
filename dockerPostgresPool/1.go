package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	envDB := os.Getenv("DATABASE_URL")
	sDb := "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	fmt.Println("envDB=", envDB)
	dbpool, err := pgxpool.New(context.Background(), sDb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	dbpool.Ping(context.Background())
	fmt.Println("about to query")

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
