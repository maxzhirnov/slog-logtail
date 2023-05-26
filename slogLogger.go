package logtail

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/slog"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	WARN           = "WARN"
	ERR            = "ERR"
	DEBUG          = "DEBUG"
)

type Logger struct {
	buf     *bytes.Buffer
	logtail *Logtail
	*slog.Logger
}

func NewLogger(logtailSecret string) Logger {
	buf := &bytes.Buffer{}
	multiWriter := io.MultiWriter(os.Stdout, buf)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	var handler slog.Handler = slog.NewJSONHandler(multiWriter, opts)
	return Logger{
		buf:     buf,
		logtail: NewLogtail(logtailSecret),
		Logger:  slog.New(handler),
	}
}

func NewLoggerWithChild(logtailSecret string, args ...any) Logger {
	buf := &bytes.Buffer{}
	multiWriter := io.MultiWriter(os.Stdout, buf)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	var handler slog.Handler = slog.NewJSONHandler(multiWriter, opts)
	logger := slog.New(handler)
	child := logger.With(args...)

	return Logger{
		buf:     buf,
		logtail: NewLogtail(logtailSecret),
		Logger:  child,
	}
}

func (l *Logger) SetDefault() {
	slog.SetDefault(l.Logger)
}

func (l *Logger) Log(logLevel LogLevel, msg string, args ...any) {

	switch logLevel {
	case INFO:
		l.Info(msg, args...)
	case WARN:
		l.Warn(msg, args...)
	case ERR:
		l.Error(msg, args...)
	case DEBUG:
		l.Debug(msg, args...)
	}

	code, err := l.logtail.Send(l.buf.String())
	l.buf.Reset()

	fmt.Println(code, err)
	if err != nil {
		return
	}
}
