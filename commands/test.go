package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
)

func TestCommand(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
  text, err := language.CurrentLanguage.Get(language.Start)
  if err != nil {
    return tg.MessageConfig{}, err
  }

  // add inline keyboard with a button
  replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
  replyMsg.ReplyToMessageID = update.Message.MessageID

  replyMsg.ReplyMarkup = tg.NewInlineKeyboardMarkup(
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonURL("Go to Google", "https://google.com"),
      tg.NewInlineKeyboardButtonData("Click me", "click_me"),
    ),
  )

  return replyMsg, nil
}
