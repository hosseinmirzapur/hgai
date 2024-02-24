package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/golangchain/utils"
)

func HandleBotUpdate(tools *utils.Tools, update *tgbotapi.Update) error {
	var err error

	// bot command
	if update.Message.IsCommand() {
		err = handleCommand(tools, update)
	}

	// image and caption
	if update.Message.Photo != nil {
		err = handleImage(tools, update)
	}

	// normal text
	if update.Message.Text != "" {
		err = handleText(tools, update)
	}

	// if not text, image or command, then the `sent` variable will be nil
	if update.Message.Text == "" && update.Message.Photo == nil {
		return fmt.Errorf("bad request: input should either be command, text or image")
	}

	// handle any error
	if err != nil {
		return err
	}

	return nil
}
