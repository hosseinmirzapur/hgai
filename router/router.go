package router

import (
	"net/http"

	"github.com/hosseinmirzapur/golangchain/handler"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/prompt", handler.SendPrompt)
	mux.HandleFunc("/health", handler.HealthCheck)
}
