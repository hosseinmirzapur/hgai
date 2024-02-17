package handler

import (
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/hosseinmirzapur/golangchain/dto"
	"github.com/hosseinmirzapur/golangchain/utils"
	"google.golang.org/api/option"
)

func SendPrompt(w http.ResponseWriter, r *http.Request) {
	err := utils.Method(r, "POST")
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	// json-decode request body
	req := dto.PromptReq{}
	body, err := req.Decode(r)
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

	var answer []genai.Part

	answer = append(answer, res.Candidates[0].Content.Parts...)

	// print out response
	utils.SuccessResponse(w, answer)
}
