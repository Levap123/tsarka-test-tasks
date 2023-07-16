package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// hardcoded because i too lazy for adding configs
const (
	host     = "postgres"
	port     = "5432"
	user     = "dev"
	password = "dev"
	dbname   = "dev"
)

func New() (*sql.DB, error) {
	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil
}
