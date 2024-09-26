package db

import "database/sql"

type DB struct {
	DbClient *sql.DB
}

/******  d3ddc269-3594-4ae9-a98d-3da1f532d48e  *******/
func NewDb() (*DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DB{DbClient: db}, nil
}
