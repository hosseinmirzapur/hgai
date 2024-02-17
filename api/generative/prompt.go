package handler

import (
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/hosseinmirzapur/golangchain/utils"
	"google.golang.org/api/option"
)

type PromptReq struct {
	Prompt string `json:"prompt,omitempty"`
}

func SendPrompt(w http.ResponseWriter, r *http.Request) {
	err := utils.Method(r, "POST")
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	// json-decode request body
	body, err := utils.DecodeSchema(r, PromptReq{})
	if err != nil {
		utils.ErrorResponse(w, "data validation error", http.StatusUnprocessableEntity)
		return
	}

	// body data validation
	if body.Prompt == "" {
		utils.ErrorResponse(w, "prompt field is required and cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	client, err := genai.NewClient(r.Context(), option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	res, err := model.GenerateContent(r.Context(), genai.Text(body.Prompt))
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// print out response
	utils.SuccessResponse(w, res)
}
