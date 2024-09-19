package main

import "github.com/bismastr/mealprep-app/internal/server"

func main() {

	server := server.NewServer()
	server.Start()

}
