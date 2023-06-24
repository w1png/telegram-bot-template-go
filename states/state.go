package states

import (
  tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type State interface {
  OnEnter(id int64, chatID int64) (tg.MessageConfig, error)
  OnExit(id int64, chatID int64) (tg.MessageConfig, error)
  OnMessage(id int64, chatID int64, message string) (tg.MessageConfig, error)

  String() string
}

