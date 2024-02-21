package main

import (
	"fmt"
	"log"
	"net/http"

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

	webhook, err := tgbot.GetWebhookInfo()
	if err != nil {
		utils.HandleError(err)
		return
	}

	log.Println(webhook)
	ch := telegram.GetUpdatesChan(tgbot)

	ai, err := gemini.New()
	if err != nil {
		utils.HandleError(err)
		return
	}
	defer ai.Client.Close()

	// serve http server concurrently with registering webhook
	go serveHttp()

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

func serveHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", nil)
}
