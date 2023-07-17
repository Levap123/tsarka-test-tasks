package hash

import "database/sql"

const hashTable = "hash_records"

type Repo struct {
	DB *sql.DB
}

func NewRepo(DB *sql.DB) *Repo {
	return &Repo{
		DB: DB,
	}
}
