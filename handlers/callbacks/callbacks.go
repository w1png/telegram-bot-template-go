package callbacks

import (
	"github.com/w1png/telegram-bot-template/types"
)

var CallbacksMap = map[string]func(data types.CallbackData) types.UpdateHandlerFunction{
	"test":            TestCallback,
	"test_command":    TestCommandCallback,
	"test_name_state": TestNameStateCallback,
}
