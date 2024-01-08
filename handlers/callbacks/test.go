package callbacks

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func TestCallback(data types.CallbackData) types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		t, err := language.LanguageInstance.Get(language.TestCallback)
		if err != nil {
			return nil, err
		}

		callback_data, err := types.NewCallbackData("test_command", nil).Marshal()
		if err != nil {
			return nil, err
		}

		back_t, err := language.LanguageInstance.Get(language.Back)
		if err != nil {
			return nil, err
		}

		return tg.NewEditMessageTextAndMarkup(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			fmt.Sprintf(t, data.Data),
			tg.NewInlineKeyboardMarkup(tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonData(back_t, string(callback_data)),
			)),
		), nil
	}
}
