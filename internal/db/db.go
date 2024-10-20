package db

import (
	"os"

	"github.com/jackc/pgx"
)

type DB struct {
	DbClient *pgx.ConnPool
}

func NewDb() (*DB, error) {
	config := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     5432,
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: "meal_prep",
		},
	}

	dbpool, err := pgx.NewConnPool(config)
	if err != nil {
		return nil, err
	}

	return &DB{DbClient: dbpool}, nil
}
