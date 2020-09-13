package server

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

const (
	initShema = `
		CREATE TABLE IF NOT EXISTS songs(
			id TEXT,
			name TEXT,
			autor TEXT,
			description TEXT
		)
	`
)

// StartDB --
func StartDB() (*sql.DB, error) {

	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	dbname := os.Getenv("APP_DB_NAME")

	pgURL := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", "localhost", 5432),
		User:   url.UserPassword(user, password),
		Path:   dbname,
	}

	options := url.Values{}
	options.Set("sslmode", "disable")

	pgURL.RawQuery = options.Encode()

	dsn := pgURL.String()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}

// InitDBSchema --
func (svr *Svr) InitDBSchema() error {
	_, err := svr.db.Exec(initShema)
	if err != nil {
		return err
	}
	return nil
}
