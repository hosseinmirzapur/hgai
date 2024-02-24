package telegram

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New() (*tgbotapi.BotAPI, error) {
	tgbot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_API_KEY"))
	if err != nil {
		return nil, err
	}

	tgbot.Debug = true // verbose logs
	return tgbot, nil
}

func GetUpdatesChan(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 // this technique is called long-polling
	updates := bot.GetUpdatesChan(u)
	return updates
}

func NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}

func DeleteCurrentWHook(bot *tgbotapi.BotAPI) (*tgbotapi.APIResponse, error) {
	return bot.MakeRequest("deleteWebhook", tgbotapi.Params{
		"drop_pending_updates": "True",
	})
}
