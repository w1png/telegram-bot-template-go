package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func StartCommand() types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		text, err := language.LanguageInstance.Get(language.Start)
		if err != nil {
			return tg.MessageConfig{}, err
		}

		replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
		replyMsg.ReplyToMessageID = update.Message.MessageID

		return replyMsg, nil
	}
}
