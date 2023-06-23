package language

import "github.com/w1png/paid-access-telegram-bot/errors"

type Russian struct {
  Values map[LanguageString]string
}

func NewRussian() *Russian {
  return &Russian{
    Values: map[LanguageString]string{
      Start: "Привет, это стартовое сообщение",
      Help: "Привет, это сообщение помощи",
      UnknownCommand: "Неизвестная команда",
      UnknownCallback: "Неизвестный коллбэк",
      UnknownError: "Неизвестная ошибка",
    },
  }
}

func (e *Russian) Get(key LanguageString) (string, error) {
  value, ok := e.Values[key]
  if !ok {
    return "", errors.NewLanguageStringError(key.String())
  }
  return value, nil
}

