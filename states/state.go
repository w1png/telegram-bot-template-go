package states

import (
	"github.com/w1png/telegram-bot-template/types"
)

type State interface {
	OnMessage() types.UpdateHandlerFunction
	OnCallback(callback_data types.CallbackData) types.UpdateHandlerFunction

	String() string
}
