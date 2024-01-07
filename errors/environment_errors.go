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
