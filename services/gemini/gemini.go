package gemini

import (
	"context"
	"errors"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type gemini struct {
	Client *genai.Client
	model  *genai.GenerativeModel
}

func New() (*gemini, error) {
	client, err := genai.NewClient(
		context.Background(),
		option.WithAPIKey(
			os.Getenv("GOOGLE_API_KEY"),
		),
	)
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-pro")

	return &gemini{
		Client: client,
		model:  model,
	}, nil
}

func (g *gemini) Generate(prompt string) ([]genai.Part, error) {
	resp, err := g.model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	if resp.Candidates == nil {
		return nil, errors.New("no candidates found")
	}

	return resp.Candidates[0].Content.Parts, nil
}
