package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/hosseinmirzapur/golangchain/api"
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

func runServer() {
	// initialize mux server
	mux := http.NewServeMux()

	// register http routes
	handler.RegisterRoutes(mux)

	// run server
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")

	log.Printf("Server started at %s:%s\n", host, port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Println("server cannot start, err:" + err.Error())
		return
	}
}
