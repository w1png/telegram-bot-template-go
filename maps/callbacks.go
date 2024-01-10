package maps

import (
	"github.com/w1png/telegram-bot-template/types"
)

var CallbackMap = CallbackMapType{
	callbacks: make(map[string]func(data types.CallbackData) types.UpdateHandlerFunction),
}

type CallbackMapType struct {
	callbacks map[string]func(data types.CallbackData) types.UpdateHandlerFunction
}

func (cm CallbackMapType) RegisterCallback(call string, callback func(data types.CallbackData) types.UpdateHandlerFunction) {
	cm.callbacks[call] = callback
}

func (cm CallbackMapType) GetCallback(call string) (func(data types.CallbackData) types.UpdateHandlerFunction, bool) {
	callback, ok := cm.callbacks[call]
	return callback, ok
}
