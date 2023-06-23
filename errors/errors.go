package errors

import "fmt"

type EnvironmentVariableError struct {
	Variable string
}

func (e *EnvironmentVariableError) Error() string {
	return fmt.Sprintf("Environment variable %s is not set", e.Variable)
}

func NewEnvironmentVariableError(variable string) *EnvironmentVariableError {
  return &EnvironmentVariableError{Variable: variable}
}


type LanguageStringError struct {
  Key string
}

func (e *LanguageStringError) Error() string {
  return fmt.Sprintf("Language string %s is not set", e.Key)
}

func NewLanguageStringError(key string) *LanguageStringError {
  return &LanguageStringError{Key: key}
}

type LanguageNotFoundError struct {
  Language string
}

func (e *LanguageNotFoundError) Error() string {
  return fmt.Sprintf("Language %s is not found", e.Language)
}

func NewLanguageNotFoundError(language string) *LanguageNotFoundError {
  return &LanguageNotFoundError{Language: language}
}

func NewLoggerNotFoundError(logger string) *LoggerNotFoundError {
  return &LoggerNotFoundError{Logger: logger}
}

type LoggerNotFoundError struct {
  Logger string
}

func (e *LoggerNotFoundError) Error() string {
  return fmt.Sprintf("Logger %s is not found", e.Logger)
}
