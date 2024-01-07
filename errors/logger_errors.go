package errors

import "fmt"

func NewLoggerNotFoundError(logger string) *LoggerNotFoundError {
	return &LoggerNotFoundError{Logger: logger}
}

type LoggerNotFoundError struct {
	Logger string
}

func (e *LoggerNotFoundError) Error() string {
	return fmt.Sprintf("Logger %s is not found", e.Logger)
}
