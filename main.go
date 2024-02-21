package main

import (
	"log"

	"github.com/hosseinmirzapur/golangchain/services/gemini"
	"github.com/hosseinmirzapur/golangchain/services/telegram"
	"github.com/hosseinmirzapur/golangchain/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	tgbot, err := telegram.New()
	if err != nil {
		utils.HandleError(err)
		return
	}
	ch := telegram.GetUpdatesChan(tgbot)

	ai, err := gemini.New()
	if err != nil {
		utils.HandleError(err)
		return
	}
	defer ai.Client.Close()

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

		log.Printf("%+v", parts)
	}
}
