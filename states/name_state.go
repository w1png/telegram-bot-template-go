package states

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NameState struct {
	Name string
}

func (s *NameState) OnEnter(id int64, chatID int64) (tg.MessageConfig, error) {
	msg := tg.NewMessage(chatID, "What is your name?")
	return msg, nil
}

func (s *NameState) OnExit(id int64, chatID int64) (tg.MessageConfig, error) {
	msg := tg.NewMessage(chatID, "Goodbye, "+s.Name+"!")
	msg.ReplyToMessageID = -1

	StateMachineInstance.RemoveState(NewStateUser(id, chatID))
	return msg, nil
}

func (s *NameState) OnMessage(id int64, chatID int64, message string) (tg.MessageConfig, error) {
	s.Name = message
	_, err := s.OnExit(id, chatID)
	if err != nil {
		return tg.MessageConfig{}, err
	}

	state := NewAgeState()
	state.(*AgeState).Name = s.Name
	StateMachineInstance.AddState(NewStateUser(chatID, chatID), state)

	return state.OnEnter(id, chatID)
}

func (s NameState) String() string {
	return "NameState"
}

func NewNameState() State {
	return &NameState{}
}
