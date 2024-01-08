package callbacks

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
)

func UnknownCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) (tgbotapi.Chattable, error) {
	text, err := language.LanguageInstance.Get(language.UnknownCallback)
	if err != nil {
		return tg.MessageConfig{}, err
	}

	message := tg.NewMessage(update.CallbackQuery.Message.Chat.ID, text)

	return message, nil
}
