package callbacks

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/states"
	"github.com/w1png/telegram-bot-template/types"
)

func TestNameStateCallback(data types.CallbackData) types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		states.StateMachineInstance.SetState(states.NewStateUser(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.Chat.ID,
		), states.NewNameState())

		text, err := language.LanguageInstance.Get(language.TestStateGetName)
		if err != nil {
			return tg.MessageConfig{}, err
		}

		return tg.NewEditMessageText(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			text,
		), nil
	}
}
