package db

import (
	"context"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
)

// DB is a struct holding a postgres connection
type DB struct {
	conn *pgx.Conn
}

// Connect returns the database connection
func Connect(dbURL string) *DB {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("unable to connect to db err: %v", err)
		os.Exit(1)
	}
	return &DB{
		conn: conn,
	}
}

//MigrateUp runs migrations
func MigrateUp(dbURL string) {
	m, _ := migrate.New(
		"file://db/migrations",
		dbURL)
	m.Up()
}
