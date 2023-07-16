package user

import "database/sql"

const userTable = "users"

type Repo struct {
	DB *sql.DB
}

func NewRepo(DB *sql.DB) *Repo {
	return &Repo{
		DB: DB,
	}
}
