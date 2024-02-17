package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

type Req struct {
	Prompt string `json:"prompt,omitempty"`
}

func SendPrompt(w http.ResponseWriter, r *http.Request) {
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
