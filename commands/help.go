package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
)

func HelpCommand(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
	text, err := language.CurrentLanguage.Get(language.Help)
	if err != nil {
		return tg.MessageConfig{}, err
	}

	replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
	replyMsg.ReplyToMessageID = update.Message.MessageID

	return replyMsg, nil
}
