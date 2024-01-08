package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func TestCommand() types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		msg := tg.NewMessage(
			update.FromChat().ID,
			"This is a test command",
		)

		msg.ReplyToMessageID = update.Message.MessageID

		rows := [][]tg.InlineKeyboardButton{}

		test_callback, err := language.LanguageInstance.Get(language.TestCallback)
		if err != nil {
			return msg, err
		}
		test_callback_data, err := types.NewCallbackData("test", "test").Marshal()
		if err != nil {
			return msg, err
		}

		rows = append(rows, tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(
				test_callback,
				string(test_callback_data),
			),
		))

		test_name_state, err := language.LanguageInstance.Get(language.TestState)
		if err != nil {
			return msg, err
		}

		test_name_state_data, err := types.NewCallbackData("test_name_state", "test_name_state").Marshal()
		if err != nil {
			return msg, err
		}

		rows = append(rows, tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(
				test_name_state,
				string(test_name_state_data),
			),
		))

		msg.ReplyMarkup = tg.NewInlineKeyboardMarkup(
			rows...,
		)

		return msg, nil
	}
}
