package logtail

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramLogger interface {
	Log(logLevel LogLevel, msgName string, body ...any)
}

type TelegramLogerObject struct {
	Token  string
	ChatID int64
	Logger *Logger
}

func NewTelegramLoger(token string, chatID int64, logger *Logger) TelegramLogger {
	return &TelegramLogerObject{
		Token:  token,
		ChatID: chatID,
		Logger: logger,
	}
}

func (tlo TelegramLogerObject) Log(logLevel LogLevel, msgName string, body ...any) {
	tlo.Logger.Log(logLevel, msgName, body...)
	bot, err := tgbotapi.NewBotAPI(tlo.Token)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = false

	msg := tgbotapi.NewMessage(tlo.ChatID, fmt.Sprintf("[[%s]]\n%s\n%s", logLevel, msgName, body))
	msg.ParseMode = "Markdown"
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
