# slog-logtail

Usage Example:
```
appInfoLogs := slog.Group(
		"app_info",
		slog.Int("pid", os.Getpid()),
		slog.String("env", appEnv),
	)

	loger := logtail.NewLoggerWithChild(LOGTAIL_SECRET, appInfoLogs)
	loger.SetDefault()

	//Telegream logger
	var tgLogger tglog.Loger = tglog.TgLoger{
		Token:  tgLoggerSecret,
		ChatID: TELEGRAM_CHAT_ID,
		Logger: &loger,
	}
	tgLogger.Log(logtail.INFO, "Starting...")
  ```
