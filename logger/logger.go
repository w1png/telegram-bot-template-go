package logger

import (
	"log"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/errors"
)

type Logger interface {
	Log(level LogLevel, message string)
	LogUpdate(update tg.Update, startTime time.Time)
}

var CurrentLogger Logger

func InitLogger(loggerType string) error {
	switch loggerType {
	case "console":
		CurrentLogger = NewConsoleLogger()
	case "no":
		CurrentLogger = NewNoLogger()
	case "":
		log.Println("No logger type specified, using no logger")
		CurrentLogger = NewNoLogger()
	default:
		return errors.NewLoggerNotFoundError(loggerType)
	}

	CurrentLogger = NewConsoleLogger()

	return nil
}
