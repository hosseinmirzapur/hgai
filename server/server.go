package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hosseinmirzapur/golangchain/router"
)

func Run() error {
	mux := http.NewServeMux()
	router.RegisterRoutes(mux)

	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	return http.ListenAndServe(addr, mux)

}
