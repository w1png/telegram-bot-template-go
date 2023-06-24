package states

import "fmt"

type StateUser struct {
  TelegramID int64
  ChatID int64
}

func NewStateUser(telegramID int64, chatID int64) StateUser {
  return StateUser{
    TelegramID: telegramID,
    ChatID: chatID,
  }
}

func (s StateUser) String() string {
  return "StateUser{TelegramID: " + fmt.Sprint(s.TelegramID) + ", ChatID: " + fmt.Sprint(s.ChatID) + "}"
}

