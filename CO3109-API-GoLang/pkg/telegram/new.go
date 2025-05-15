package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TeleBot interface {
	SendMessage(chatID int64, text string) (tgbotapi.Message, error)
}

type implTeleBot struct {
	bot *tgbotapi.BotAPI
}

func NewManager(teleBotKey string) TeleBot {
	bot, err := tgbotapi.NewBotAPI(teleBotKey)
	if err != nil {
		panic(err)
	}
	return &implTeleBot{
		bot: bot,
	}
}
