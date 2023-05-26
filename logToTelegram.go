package logtail

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgLoger interface {
	Log(logLevel LogLevel, msgName string, body ...any)
}

type tgLoger struct {
	Token  string
	ChatID int64
	Logger *Logger
}

func (l tgLoger) Log(logLevel LogLevel, msgName string, body ...any) {
	go func() { l.Logger.Log(logLevel, msgName, body...) }()
	bot, err := tgbotapi.NewBotAPI(l.Token)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = false

	msg := tgbotapi.NewMessage(l.ChatID, fmt.Sprintf("[[%s]]\n%s\n%s", logLevel, msgName, body))
	msg.ParseMode = "Markdown"
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
