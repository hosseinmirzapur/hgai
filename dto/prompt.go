package dto

import (
	"encoding/json"
	"net/http"
)

type PromptReq struct {
	Prompt string `json:"prompt,omitempty"`
}

func (req *PromptReq) Decode(r *http.Request) (*PromptReq, error) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
