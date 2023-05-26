package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	l "github.com/maxzhirnov/slog-logtail"
	"golang.org/x/exp/slog"
)

func init() {
	err := godotenv.Load(".env")
	// err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

}

func main() {
	var (
		TG_LOGER_SECRET = os.Getenv("TG_LOGER_SECRET")
		LOGTAIL_SECRET  = os.Getenv("LOGTAIL_SECRET")
		DEBUG_CHAT_ID   = os.Getenv("DEBUG_CHAT_ID")
		APP_ENV         = "development"
	)
	appInfoLogs := slog.Group(
		"app_info",
		slog.Int("pid", os.Getpid()),
		slog.String("env", APP_ENV),
	)
	logger := l.NewLoggerWithChild(LOGTAIL_SECRET, appInfoLogs)
	logger.SetDefault()

	//Telegream logger
	chatId, err := strconv.ParseInt(DEBUG_CHAT_ID, 10, 64)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	var telegramLogger l.TelegramLogger = l.NewTelegramLoger(TG_LOGER_SECRET, chatId, &logger)
	telegramLogger.Log(l.INFO, "Test message 1...")
	telegramLogger.Log(l.INFO, "Test message 2...")
	telegramLogger.Log(l.INFO, "Test message 3...")

	telegramLogger.Log(l.WARN, "Test Message with body",
		slog.String("foo", "bar"),
		slog.Int("baz", 42),
		slog.Bool("qux", true),
	)

	telegramLogger.Log(l.WARN, "Test Message with group body",
		slog.Group("foo",
			slog.String("foo", "bar"),
			slog.Int("baz", 42),
			slog.Bool("qux", true)),
	)
}
