package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/golangchain/services/gemini"
)

type Tools struct {
	AI  *gemini.Gemini
	Bot *tgbotapi.BotAPI
}
