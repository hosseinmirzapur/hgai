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
}

// @todo add commands in botfather
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
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		"Soon...",
	)
	_, err := tools.Bot.Send(msg)
	return err
}

func handleBalance(tools *utils.Tools, update *tgbotapi.Update) error {
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		"Soon...",
	)
	_, err := tools.Bot.Send(msg)
	return err
}

func handleHelp(tools *utils.Tools, update *tgbotapi.Update) error {
	cmds, err := tools.Bot.GetMyCommands()
	if err != nil {
		return err
	}

	show := "available\n"

	for _, cmd := range cmds {
		show += fmt.Sprintf("%s: %s\n", cmd.Command, cmd.Description)
	}

	_, err = tools.Bot.Send(
		telegram.NewMessage(update.Message.Chat.ID, show),
	)
	return err

}

func handleReferral(tools *utils.Tools, update *tgbotapi.Update) error {
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		"Soon...",
	)
	_, err := tools.Bot.Send(msg)
	return err
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
