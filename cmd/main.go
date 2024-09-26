package main

import (
	"log"

	"github.com/bismastr/mealprep-app/internal/db"
	"github.com/bismastr/mealprep-app/internal/server"
)

func main() {

	server := server.NewServer()
	server.Start()
	db, _ := db.NewDb()

	err := db.DbClient.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
