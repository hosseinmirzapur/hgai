package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	// run http server
	runServer()
}
