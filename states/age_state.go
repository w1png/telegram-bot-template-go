package states

import (
	"fmt"
	"strconv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AgeState struct {
	Name string
	Age  string
}

func (s *AgeState) OnEnter(id int64, chatID int64) (tg.MessageConfig, error) {
	msg := tg.NewMessage(chatID, fmt.Sprintf("Ok, %s. What is your age?", s.Name))
	return msg, nil
}

func (s *AgeState) OnExit(id int64, chatID int64) (tg.MessageConfig, error) {
	msg := tg.NewMessage(chatID, "Goodbye, "+s.Name+"!")
	msg.ReplyToMessageID = -1

	StateMachineInstance.RemoveState(NewStateUser(id, chatID))
	return msg, nil
}

func (s *AgeState) OnMessage(id int64, chatID int64, message string) (tg.MessageConfig, error) {
	age, err := strconv.Atoi(message)
	if err != nil {
		return tg.MessageConfig{}, err
	}

	m := ""
	if age < 18 {
		m = fmt.Sprintf("You are not allowed to enter, %s", s.Name)
	} else {
		m = fmt.Sprintf("You are allowed to enter, %s", s.Name)
	}

	msg := tg.NewMessage(chatID, m)
	_, err = s.OnExit(id, chatID)

	return msg, err
}

func (s AgeState) String() string {
	return "AgeState"
}

func NewAgeState() State {
	return &AgeState{}
}
