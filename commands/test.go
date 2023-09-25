package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/states"
)

func TestCommand(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
	state := states.NewNameState()

	states.StateMachineInstance.AddState(states.NewStateUser(msg.From.ID, msg.Chat.ID), state)

	return state.OnEnter(msg.From.ID, msg.Chat.ID)
}
