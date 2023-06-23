package logger

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  "time"
)

type NoLogger struct {}

func NewNoLogger() *NoLogger {
  return &NoLogger{}
}

func (l *NoLogger) Log(level LogLevel, message string) {}

func (l *NoLogger) LogUpdate(update tg.Update, time time.Time) {}
