package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ConsoleLogger struct {}

func NewConsoleLogger() *ConsoleLogger {
  return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(level LogLevel, message string) {
  log.Printf("[%s] %s\n", level.String(), message)
  if level == Fatal {
    os.Exit(1)
  }
}

func (l *ConsoleLogger) LogUpdate(update tg.Update, startTime time.Time) {
  username := "unknown username"
  text := "unknown text or data"

  if update.Message != nil {
    username = update.Message.From.UserName
    text = fmt.Sprintf("Message text: %s", update.Message.Text)
  } else if update.CallbackQuery != nil {
    username = update.CallbackQuery.From.UserName
    text = fmt.Sprintf("Callback data: %s", update.CallbackQuery.Data)
  }

  log.Printf("[%s] Update: [From: %s] [Data: %s] [Took: %s]\n", Info.String(), username, text, time.Since(startTime))
}

