package main

import (
	"github.com/hosseinmirzapur/golangchain/config"
	"github.com/hosseinmirzapur/golangchain/server"
	"github.com/hosseinmirzapur/golangchain/services/gemini"
	"github.com/hosseinmirzapur/golangchain/services/telegram"
	"github.com/hosseinmirzapur/golangchain/utils"
)

func main() {
	// load .env variables
	config.LoadEnv()

	// start telegram bot
	bot, err := telegram.New()
	if err != nil {
		utils.HandleError(err)
		return
	}

	// establish gemini client connection
	ai, err := gemini.New()
	if err != nil {
		utils.HandleError(err)
		return
	}
	defer ai.Client.Close()

	tools := utils.Tools{
		AI:  ai,
		Bot: bot,
	}

	// serve the bot
	server.ListenForUpdates(&tools)
}
