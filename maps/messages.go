package maps

import (
	"github.com/w1png/telegram-bot-template/types"
)

var MessagesMap = MessagesMapType{
	messages: make(map[string]func() types.UpdateHandlerFunction),
}

type MessagesMapType struct {
	messages map[string]func() types.UpdateHandlerFunction
}

func (m MessagesMapType) ResgisterMessage(call string, message func() types.UpdateHandlerFunction) {
	m.messages[call] = message
}

func (m MessagesMapType) GetMessage(call string) (func() types.UpdateHandlerFunction, bool) {
	message, ok := m.messages[call]
	return message, ok
}
