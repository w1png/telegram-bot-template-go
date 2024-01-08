package language

type LanguageString string

const (
	Start           LanguageString = "Start"
	Help            LanguageString = "Help"
	UnknownCommand  LanguageString = "UnknownCommand"
	UnknownCallback LanguageString = "UnknownCallback"
	UnknownError    LanguageString = "UnknownError"
	Back            LanguageString = "Back"

	TestCommand       LanguageString = "TestCommand"
	TestCallback      LanguageString = "TestCallback"
	TestState         LanguageString = "TestStateCallback"
	TestStateGetName  LanguageString = "TestState"
	TestStateGreeting LanguageString = "TestStateGreeting"
)
