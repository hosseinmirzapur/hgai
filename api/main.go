package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
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

type Req struct {
	Prompt string `json:"prompt,omitempty"`
}

func baseFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fine!"))
}

func sendPrompt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handleErr(w, errors.New("method not allowed"))
		return
	}

	// json-decode request body
	body, err := decodeBody(r)
	if err != nil {
		handleErr(w, errors.New("bad request structure"))
		return
	}

	// body data validation
	if body.Prompt == "" {
		handleErr(w, errors.New("'prompt' field is required and cannot be empty"))
		return
	}

	// load google-ai model
	model, err := googleai.New(r.Context(), googleai.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		handleErr(w, err)
		return
	}

	// generate response from text input
	res, err := llms.GenerateFromSinglePrompt(r.Context(), model, body.Prompt)
	if err != nil {
		handleErr(w, err)
		return
	}

	// print out response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func decodeBody(r *http.Request) (*Req, error) {
	var req Req

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func handleErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
