package callbacks

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/handlers/commands"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func TestCommandCallback(data types.CallbackData) types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		update.Message = update.CallbackQuery.Message

		chattable, err := commands.TestCommand()(bot, update, user)
		if err != nil {
			return nil, err
		}

		chattable_msg := chattable.(tg.MessageConfig)
		msg := tg.NewEditMessageTextAndMarkup(
			chattable_msg.ChatID,
			update.CallbackQuery.Message.MessageID,
			chattable_msg.Text,
			chattable_msg.ReplyMarkup.(tg.InlineKeyboardMarkup),
		)

		return msg, nil
	}
}
