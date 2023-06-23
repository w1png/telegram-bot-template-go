package callbacks

import (
  tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  "github.com/w1png/telegram-bot-template/language"
)

func UnknownCallback(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
  text, err := language.CurrentLanguage.Get(language.UnknownCallback)
  if err != nil {
    return tg.MessageConfig{}, err
  }

  message := tg.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
  
  return message, nil
}

