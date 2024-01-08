package errors

import "fmt"

type MessageSendError struct {
	Err string
}

func NewMessageSendError(err string) *MessageSendError {
	return &MessageSendError{Err: err}
}

func (e *MessageSendError) Error() string {
	return fmt.Sprintf("Message send error: %s", e.Err)
}
