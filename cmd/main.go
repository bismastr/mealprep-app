package main

import (
	"fmt"

	"github.com/bismastr/mealprep-app/internal/db"
	"github.com/bismastr/mealprep-app/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server := server.NewServer()

	db, err := db.NewDb()
	if err != nil {
		fmt.Println(err)
	}

	defer db.DbClient.Close()
	server.Start()
}
