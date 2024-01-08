package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func HelloMessage() types.UpdateHandlerFunction {
	return func(bot *tgbotapi.BotAPI, update tgbotapi.Update, user *models.User) (tgbotapi.Chattable, error) {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "Hello"), nil
	}
}
