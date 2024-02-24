package handler

import (
	"fmt"
	"slices"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/golangchain/services/telegram"
	"github.com/hosseinmirzapur/golangchain/utils"
)

var commands = []string{
	"start",
	"pay",
	"balance",
	"help",
	"referral",
	"clean",
	"stop",
}

func handleCommand(tools *utils.Tools, update *tgbotapi.Update) error {
	var err error

	switch update.Message.Command() {
	case "start":
		err = handleStart(tools, update)
	case "pay":
		err = handlePay(tools, update)
	case "balance":
		err = handleBalance(tools, update)
	case "help":
		err = handleHelp(tools, update)
	case "referral":
		err = handleReferral(tools, update)
	case "clean":
		err = handleClean(tools, update)
	case "stop":
		err = handleStop(tools, update)
	default:
		err = handleDefault(tools, update)
	}

	return err
}

func handleStart(tools *utils.Tools, update *tgbotapi.Update) error {
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("Welcome back dear @%s!", update.Message.Chat.UserName),
	)

	_, err := tools.Bot.Send(msg)
	return err
}

func handlePay(tools *utils.Tools, update *tgbotapi.Update) error {
	return nil
}

func handleBalance(tools *utils.Tools, update *tgbotapi.Update) error {
	return nil
}

func handleHelp(tools *utils.Tools, update *tgbotapi.Update) error {
	return nil
}

func handleReferral(tools *utils.Tools, update *tgbotapi.Update) error {
	return nil
}

func handleClean(tools *utils.Tools, update *tgbotapi.Update) error {
	return nil
}

func handleStop(tools *utils.Tools, update *tgbotapi.Update) error {
	return nil
}

func handleDefault(tools *utils.Tools, update *tgbotapi.Update) error {
	if !slices.Contains(commands, update.Message.Command()) {
		err := fmt.Errorf("this command is not supported by the bot")
		msg := telegram.NewMessage(update.Message.Chat.ID, err.Error())

		_, err = tools.Bot.Send(msg)

		return err
	}
	return nil
}
