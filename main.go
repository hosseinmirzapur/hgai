package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	// remove any associated webhook before establishing a new one
	// this is the rawest form of calling the telegram API
	// the `DeleteWebhook` method was not included in the library
	response, err := tgbot.MakeRequest("", tgbotapi.Params{
		"deleteWebhook": "True",
	})
	if err != nil {
		utils.HandleError(err)
		return
	}
	if !response.Ok {
		utils.HandleError(fmt.Errorf("failed to delete webhook: %+v", response))
		return
	}

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
