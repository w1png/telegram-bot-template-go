package errors

import "fmt"

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
