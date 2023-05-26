package logtail

import (
	"encoding/json"
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
	msgText := tlo.Logger.Log(logLevel, msgName, body...)
	var data interface{}
	json.Unmarshal([]byte(msgText), &data)
	formattedMsg := formatMessage(data, 0)
	fmt.Println(formattedMsg)

	bot, err := tgbotapi.NewBotAPI(tlo.Token)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = false

	msg := tgbotapi.NewMessage(tlo.ChatID, formattedMsg)
	// msg.ParseMode = "Markdown"
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
