package main

import (
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/errors"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/logger"
	"github.com/w1png/telegram-bot-template/maps"
	"github.com/w1png/telegram-bot-template/states"
	"github.com/w1png/telegram-bot-template/storage"
	"github.com/w1png/telegram-bot-template/types"
)

type Bot struct {
	Bot     *tg.BotAPI
	timeout int
}

func NewBot(timeout int) (*Bot, error) {
	bot, err := tg.NewBotAPI(config.ConfigInstance.TelegramToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = config.ConfigInstance.IsDebug
	return &Bot{Bot: bot, timeout: timeout}, nil
}

func (b *Bot) Run() error {
	u := tg.NewUpdate(0)
	u.Timeout = b.timeout

	updates := b.Bot.GetUpdatesChan(u)
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		go func(update tg.Update) {
			startTime := time.Now()
			b.HandleUpdate(update)
			logger.LoggerInstance.LogUpdate(update, startTime)
		}(update)
	}

	return nil
}

func (b *Bot) Stop() {
	b.Bot.StopReceivingUpdates()
}

func (b *Bot) SendUnknownError(chat_id int64) {
	s, err := language.LanguageInstance.Get(language.UnknownError)
	if err != nil {
		logger.LoggerInstance.Log(logger.Fatal, err.Error())
		return
	}
	if _, err = b.Bot.Send(tg.NewMessage(chat_id, s)); err != nil {
		logger.LoggerInstance.Log(logger.Error, errors.NewMessageSendError(err.Error()).Error())
		return
	}
}

func (b *Bot) SendUnknownActionError(chat_id int64) {
	s, err := language.LanguageInstance.Get(language.UnknownCommand)
	if err != nil {
		logger.LoggerInstance.Log(logger.Fatal, err.Error())
		return
	}
	if _, err = b.Bot.Send(tg.NewMessage(chat_id, s)); err != nil {
		logger.LoggerInstance.Log(logger.Error, errors.NewMessageSendError(err.Error()).Error())
		return
	}
}

func (b *Bot) HandleUpdate(update tg.Update) {
	var telegram_id int64
	if update.Message != nil {
		telegram_id = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		telegram_id = update.CallbackQuery.From.ID
	}
	if err := storage.StorageInstance.CreateUserIfDoesntExist(telegram_id); err != nil {
		logger.LoggerInstance.Log(logger.Error, err.Error())
		b.SendUnknownError(update.Message.Chat.ID)
		return
	}

	user, err := storage.StorageInstance.GetUserByTelegramID(telegram_id)
	if err != nil {
		b.SendUnknownError(update.Message.Chat.ID)
		logger.LoggerInstance.Log(logger.Error, err.Error())
		return
	}

	state, haveState := states.StateMachineInstance.GetState(
		states.NewStateUser(telegram_id, telegram_id),
	)

	var f types.UpdateHandlerFunction

	if update.CallbackQuery != nil {
		callback_data, err := types.UnmarshalCallbackData(update.CallbackQuery.Data)
		if err != nil {
			logger.LoggerInstance.Log(logger.Error, err.Error())
			return
		}

		if haveState {
			f = state.OnCallback(*callback_data)
		} else {
			callback_f, ok := maps.CallbackMap.GetCallback(callback_data.Call)
			if !ok {
				logger.LoggerInstance.Log(logger.Error, errors.NewUnknownCallbackError(callback_data.Call).Error())
				b.SendUnknownActionError(update.CallbackQuery.From.ID)
				return
			}
			f = callback_f(*callback_data)
		}

	} else if update.Message != nil && update.Message.IsCommand() {
		command_f, ok := maps.CommandsMap.GetCommand(update.Message.Command())
		if !ok {
			b.SendUnknownActionError(update.Message.Chat.ID)
			return
		}

		states.StateMachineInstance.DeleteState(states.NewStateUser(telegram_id, telegram_id))

		f = command_f()
	} else if update.Message != nil && !update.Message.IsCommand() {
		if haveState {
			f = state.OnMessage()
		} else {
			message_f, ok := maps.MessagesMap.GetMessage(update.Message.Text)
			if !ok {
				b.SendUnknownActionError(update.Message.Chat.ID)
				return
			}
			f = message_f()
		}

	}

	chattable, err := f(b.Bot, update, user)
	if err != nil {
		b.SendUnknownError(update.Message.Chat.ID)
		logger.LoggerInstance.Log(logger.Error, err.Error())
		return
	}

	if _, err := b.Bot.Send(chattable); err != nil {
		logger.LoggerInstance.Log(logger.Error, err.Error())
		return
	}
}
