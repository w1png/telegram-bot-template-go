package types

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/models"
)

type UpdateHandlerFunction func(bot *tgbotapi.BotAPI, update tgbotapi.Update, user *models.User) (tgbotapi.Chattable, error)
