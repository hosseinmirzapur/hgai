package main

import (
	"github.com/hosseinmirzapur/golangchain/services/bot"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bot.SendToBot()
}
