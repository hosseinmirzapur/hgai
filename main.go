package main

import (
	"log"

	"github.com/hosseinmirzapur/golangchain/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("environment cannot be loaded")
		return
	}

	err = server.Run()
	if err != nil {
		log.Println("cannot start the server")
		return
	}
}
