package language

import (
	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/errors"
)

type Language interface {
	Get(key LanguageString) (string, error)
}

var LanguageInstance Language

func setLanguage(language Language) {
	LanguageInstance = language
}

func InitLanguage() error {
	switch config.ConfigInstance.Language {
	case "en":
		setLanguage(NewEnglish())
	case "ru":
		setLanguage(NewRussian())
	default:
		return errors.NewLanguageNotFoundError(config.ConfigInstance.Language)
	}
	return nil
}
