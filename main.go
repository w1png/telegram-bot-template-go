package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/logger"
	"github.com/w1png/telegram-bot-template/states"
	"github.com/w1png/telegram-bot-template/storage"
)

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
