package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/handlers/callbacks"
	"github.com/w1png/telegram-bot-template/handlers/commands"
	"github.com/w1png/telegram-bot-template/handlers/messages"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/logger"
	"github.com/w1png/telegram-bot-template/maps"
	"github.com/w1png/telegram-bot-template/states"
	"github.com/w1png/telegram-bot-template/storage"
)

func initCommandsMap() {
	maps.CommandsMap.RegisterCommand("start", commands.StartCommand)
	maps.CommandsMap.RegisterCommand("test", commands.TestCommand)
}

func initMessagesMap() {
	maps.MessagesMap.ResgisterMessage("hello", messages.HelloMessage)
}

func initCallbacksMap() {
	maps.CallbackMap.RegisterCallback("test", callbacks.TestCallback)
	maps.CallbackMap.RegisterCallback("test_command", callbacks.TestCommandCallback)
	maps.CallbackMap.RegisterCallback("test_name_state", callbacks.TestNameStateCallback)
}

func main() {
	var err error
	if err = config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	states.InitStateMachine()

	if err = logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	err = storage.InitStorage()
	if err != nil {
		logger.LoggerInstance.Log(logger.Fatal, err.Error())
	}

	err = language.InitLanguage()
	if err != nil {
		logger.LoggerInstance.Log(logger.Fatal, err.Error())
	}

	initCommandsMap()
	initMessagesMap()
	initCallbacksMap()

	bot, err := NewBot(60)
	if err != nil {
		logger.LoggerInstance.Log(logger.Fatal, err.Error())
	}

	log.Printf("Bot started as @%v\n", bot.Bot.Self.UserName)
	go func() {
		if err := bot.Run(); err != nil {
			logger.LoggerInstance.Log(logger.Fatal, err.Error())
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	<-ctx.Done()
	fmt.Printf("Shutting down bot\n")
	bot.Stop()
}
