package telegram

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_API_KEY"))
	if err != nil {
		return nil, err
	}

	bot.Debug = true // verbose logs

	return bot, nil

}

func GetUpdatesChan(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 // seconds
	updates := bot.GetUpdatesChan(u)
	return updates
}

func UserMessage(received tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(received.Message.Chat.ID, received.Message.Text)
}

func NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}
