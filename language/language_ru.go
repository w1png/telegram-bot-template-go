package language

import "github.com/w1png/telegram-bot-template/errors"

type Russian struct {
	Values map[LanguageString]string
}

func NewRussian() *Russian {
	return &Russian{
		Values: map[LanguageString]string{
			Start:          "Привет, это стартовое сообщение",
			Help:           "Привет, это сообщение помощи",
			UnknownCommand: "Неизвестная команда",
			UnknownError:   "Неизвестная ошибка",
			Back:           "Назад",

			TestCommand:       "Это тестовая команда. Нажмите на кнопки ниже для проверки состояний и обратной связи",
			TestCallback:      "Это тестовый колбэк: \"%+v\"",
			TestState:         "Это тестовое состояние",
			TestStateGetName:  "Пожалуйста, введите свое имя:",
			TestStateGreeting: "Привет, %s!",
		},
	}
}

func (e *Russian) Get(key LanguageString) (string, error) {
	value, ok := e.Values[key]
	if !ok {
		return "", errors.NewLanguageStringError(string(key))
	}
	return value, nil
}
