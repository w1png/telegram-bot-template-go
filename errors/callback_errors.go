package errors

import "fmt"

type UnmarshalCallbackError struct {
	Err error
}

func NewUnmarshalCallbackError(err error) *UnmarshalCallbackError {
	return &UnmarshalCallbackError{Err: err}
}

func (u *UnmarshalCallbackError) Error() string {
	return u.Err.Error()
}

type UnknownCallbackError struct {
	Callback string
}

func NewUnknownCallbackError(callback string) *UnknownCallbackError {
	return &UnknownCallbackError{Callback: callback}
}

func (u *UnknownCallbackError) Error() string {
	return fmt.Sprintf("Unknown callback %s", u.Callback)
}
