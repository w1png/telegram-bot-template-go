package states

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

type NameState struct {
	Name string
}

func (s *NameState) OnMessage() types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		text, err := language.LanguageInstance.Get(language.TestStateGreeting)
		if err != nil {
			return tg.MessageConfig{}, err
		}

		StateMachineInstance.DeleteState(NewStateUser(update.Message.From.ID, update.Message.Chat.ID))

		return tg.NewMessage(update.Message.Chat.ID, fmt.Sprintf(text, update.Message.Text)), nil
	}
}

func (s *NameState) OnCallback(callbackData types.CallbackData) types.UpdateHandlerFunction {
	return nil
}

func (s NameState) String() string {
	return "NameState"
}

func NewNameState() State {
	return &NameState{}
}
