//go:build integration
// +build integration

package create_group_test

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectionPostgres() *pgxpool.Pool {
	url := "postgres://root:root@localhost:5432/store_db?sslmode=disable"
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}

func TruncateTable(conn *pgxpool.Pool) {
	defer conn.Close()
	conn.Exec(context.Background(), "TRUNCATE TABLE groups RESTART IDENTITY CASCADE;")
}
