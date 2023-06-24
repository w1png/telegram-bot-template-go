package main

import (
	"log"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/callbacks"
	"github.com/w1png/telegram-bot-template/commands"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/logger"
	"github.com/w1png/telegram-bot-template/messages"
	"github.com/w1png/telegram-bot-template/states"
)

type Bot struct {
  Bot *tg.BotAPI
  timeout int
}

func NewBot(token string, timeout int, debug bool) (*Bot, error) {
  bot, err := tg.NewBotAPI(token)
  if err != nil {
    return nil, err
  }
  bot.Debug = debug
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
      b.HandleUpdate(update)
    }(update)
  }

  return nil
}

func (b *Bot) Stop() {
  b.Bot.StopReceivingUpdates()
}

func (b *Bot) HandleUpdate(update tg.Update) {
  startTime := time.Now()

  var msg tg.MessageConfig
  var err error
  var shouldEdit bool
  var editMessage tg.Message

  // callbacks
  if update.CallbackQuery != nil {
    switch update.CallbackQuery.Data {
    default:
      msg, err = callbacks.UnknownCallback(update.Message, update)
      shouldEdit = true
    }

    editMessage = *update.CallbackQuery.Message
  }

  // commands
  if update.Message != nil && update.Message.IsCommand() {
    switch update.Message.Command() {
    case "start":
      msg, err = commands.StartCommand(update.Message, update)
    case "help":
      msg, err = commands.HelpCommand(update.Message, update)
    case "test":
      msg, err = commands.TestCommand(update.Message, update)
    default:
      msg, err = commands.UnknownCommand(update.Message, update)
    }

    editMessage = *update.Message
  }

  if update.Message != nil && !update.Message.IsCommand() {
    currentState, ok := states.StateMachineInstance.States[states.NewStateUser(
      update.Message.Chat.ID,
      update.Message.From.ID),
    ]

    if ok {
      msg, err = currentState.OnMessage(update.Message.From.ID, update.Message.Chat.ID, update.Message.Text)
    } else {
      switch update.Message.Text {
      default:
        msg, err = messages.UnknownMessage(update.Message, update)
      }
    }
  }

  // send or edit the msg
  if err == nil {
    if msg.ReplyToMessageID == -1 {
      return
    }

    if shouldEdit {
      markup := tg.NewInlineKeyboardMarkup([]tg.InlineKeyboardButton{})
      if msg.ReplyMarkup != nil {
        markup = msg.ReplyMarkup.(tg.InlineKeyboardMarkup)
      }

      if _, err = b.Bot.Send(tg.NewEditMessageTextAndMarkup(
        editMessage.Chat.ID,
        editMessage.MessageID,
        msg.Text,
        markup,
      )); err != nil {
        logger.CurrentLogger.Log(logger.Error, err.Error())
      }
    } else {
      if _, err = b.Bot.Send(msg); err != nil {
        log.Println(err)
      }
    }
  }

  // if error occured during callback or command processing
  if err != nil {
    text, _ := language.CurrentLanguage.Get(language.UnknownError)
    msg = tg.NewMessage(update.Message.Chat.ID, text)
    msg.ReplyToMessageID = update.Message.MessageID
    
    if _, err = b.Bot.Send(msg); err != nil {
      logger.CurrentLogger.Log(logger.Error, err.Error())
    }

    logger.CurrentLogger.Log(logger.Error, err.Error())
  }

  logger.CurrentLogger.LogUpdate(update, startTime)
}

