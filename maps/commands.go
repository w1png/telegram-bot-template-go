package maps

import (
	"github.com/w1png/telegram-bot-template/types"
)

var CommandsMap = CommandsMapType{
	commands: make(map[string]func() types.UpdateHandlerFunction),
}

type CommandsMapType struct {
	commands map[string]func() types.UpdateHandlerFunction
}

func (cm CommandsMapType) RegisterCommand(call string, command func() types.UpdateHandlerFunction) {
	cm.commands[call] = command
}

func (cm CommandsMapType) GetCommand(call string) (func() types.UpdateHandlerFunction, bool) {
	command, ok := cm.commands[call]
	return command, ok
}
