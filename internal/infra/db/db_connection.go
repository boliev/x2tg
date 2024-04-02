package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

type DB struct {
}

func NewDBConnection(host string, port int, user string, password string, dbName string) (*sql.DB, error) {
	var db *sql.DB
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("DB connection success")

	return db, nil
}
