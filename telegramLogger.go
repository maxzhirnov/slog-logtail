package logtail

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramLoger struct {
	Token  string
	ChatID int64
	Logger *Logger
}

func NewTelegramLoger(token string, chatID int64, logger *Logger) *TelegramLoger {
	return &TelegramLoger{
		Token:  token,
		ChatID: chatID,
		Logger: logger,
	}
}

func (tl TelegramLoger) Log(logLevel LogLevel, msgName string, body ...any) {
	go func() { tl.Logger.Log(logLevel, msgName, body...) }()
	bot, err := tgbotapi.NewBotAPI(tl.Token)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = false

	msg := tgbotapi.NewMessage(tl.ChatID, fmt.Sprintf("[[%s]]\n%s\n%s", logLevel, msgName, body))
	msg.ParseMode = "Markdown"
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
