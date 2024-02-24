package main

import (
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/hosseinmirzapur/golangchain/pkg"
)

func main() {
	client := pkg.InitModels()
	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)
	pkg.StartBot()
}
