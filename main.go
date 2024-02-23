package main

import (
	"fmt"
	"log"

	"github.com/hosseinmirzapur/golangchain/services/gemini"
	"github.com/hosseinmirzapur/golangchain/services/telegram"
	"github.com/hosseinmirzapur/golangchain/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// initialize telegram bot
	tgbot, err := telegram.New()
	if err != nil {
		utils.HandleError(err)
		return
	}

	// get current bot webhook information
	webhook, err := tgbot.GetWebhookInfo()
	if err != nil {
		utils.HandleError(err)
		return
	}
	log.Println(webhook)

	// initialize google gemini
	ai, err := gemini.New()
	if err != nil {
		utils.HandleError(err)
		return
	}
	defer ai.Client.Close()

	// listen for any bot update / read from channel
	ch := telegram.GetUpdatesChan(tgbot)

	for received := range ch {
		if received.Message == nil {
			continue
		}

		// get user's sent message
		sent := telegram.UserMessage(received)

		// set reply to the user
		sent.ReplyToMessageID = received.Message.MessageID

		// AI's response
		parts, err := ai.Generate(sent.Text)
		if err != nil {
			utils.HandleError(err)
			tgbot.Send(
				telegram.NewMessage(received.Message.Chat.ID, err.Error()),
			)
		}

		tgbot.Send(
			telegram.NewMessage(
				received.Message.Chat.ID,
				fmt.Sprintf("%+v", parts[0]),
			),
		)
	}
}
