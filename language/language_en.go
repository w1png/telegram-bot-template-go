package language

import "github.com/w1png/telegram-bot-template/errors"

type English struct {
	Values map[LanguageString]string
}

func NewEnglish() *English {
	return &English{
		Values: map[LanguageString]string{
			Start:          "Hello, this is a start message",
			Help:           "Hello, this is a help message",
			UnknownCommand: "Unknown command",
			UnknownError:   "Unknown error",
			Back:           "Back",

			TestCommand:       "This is a test command. Press the buttons below to test states and callbacks",
			TestCallback:      "Test callback",
			TestState:         "Test state: \"%+v\"",
			TestStateGetName:  "Please enter your name:",
			TestStateGreeting: "Hello, %s!",
		},
	}
}

func (e *English) Get(key LanguageString) (string, error) {
	value, ok := e.Values[key]
	if !ok {
		return "", errors.NewLanguageStringError(string(key))
	}
	return value, nil
}
