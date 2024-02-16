package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func runServer() {
	// initialize mux server
	mux := http.NewServeMux()

	// register http routes
	registerRoutes(mux)

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

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", baseFunc)
	mux.HandleFunc("/prompt", sendPrompt)
}
