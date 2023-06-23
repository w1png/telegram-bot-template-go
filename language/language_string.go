package language

type LanguageString int

const (
  Start LanguageString = iota
  Help
  UnknownCommand
  UnknownCallback
  UnknownError
)

func (l LanguageString) String() string {
  return [...]string{
    "Start",
    "Help",
    "UnknownCommand",
    "UnknownCallback",
    "UnknownError",
  }[l]
}

