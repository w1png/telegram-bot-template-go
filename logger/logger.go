package logger

import (
	"log"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/errors"
)

type Logger interface {
	Log(level LogLevel, message string)
	LogUpdate(update tg.Update, startTime time.Time)
}

var LoggerInstance Logger

func InitLogger() error {
	switch config.ConfigInstance.LoggerType {
	case "console":
		LoggerInstance = NewConsoleLogger()
	case "no":
		LoggerInstance = NewNoLogger()
	case "":
		log.Println("No logger type specified, using no logger")
		LoggerInstance = NewNoLogger()
	default:
		return errors.NewLoggerNotFoundError(config.ConfigInstance.LoggerType)
	}

	LoggerInstance = NewConsoleLogger()

	return nil
}
