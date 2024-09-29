package main

import (
	"fmt"

	"github.com/bismastr/mealprep-app/internal/db"
	"github.com/bismastr/mealprep-app/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	//Load .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	//Init Db
	db, err := db.NewDb()
	if err != nil {
		fmt.Println(err)
	}
	defer db.DbClient.Close()

	server := server.NewServer(db)
	server.Start()

}
