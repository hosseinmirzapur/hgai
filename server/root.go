package server

import (
	"log"

	"github.com/hosseinmirzapur/golangchain/handler"
	"github.com/hosseinmirzapur/golangchain/services/telegram"
	"github.com/hosseinmirzapur/golangchain/utils"
)

func ListenForUpdates(tools *utils.Tools) {

	// remove any associated webhook before establishing new one
	// this is the rawest form of calling the telegram API
	// the DeleteWebhook method was not included in the library
	response, err := telegram.DeleteCurrentWHook(tools.Bot)
	if err != nil {
		utils.HandleError(err)
		return
	}
	if !response.Ok {
		log.Println("unable to remove webhook, either it didn't exist or operation failed.")
	}

	// listen for any bot update / read from channel
	ch := telegram.GetUpdatesChan(tools.Bot)

	for received := range ch {
		if received.Message == nil {
			continue
		}

		handler.HandleBotUpdate(tools, &received)

	}
}
