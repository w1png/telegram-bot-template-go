package logger

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
	Fatal
)

func (l LogLevel) String() string {
	return [...]string{
		"Debug",
		"Info",
		"Warning",
		"Error",
		"Fatal",
	}[l]
}
