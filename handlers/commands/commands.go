package commands

import "github.com/w1png/telegram-bot-template/types"

var CommandsMap = map[string]func() types.UpdateHandlerFunction{
	"start": StartCommand,
	"test":  TestCommand,
}
