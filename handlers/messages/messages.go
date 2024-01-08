package messages

import (
	"github.com/w1png/telegram-bot-template/types"
)

var MessagesMap = map[string]func() types.UpdateHandlerFunction{
	"Hello": HelloMessage,
}
