package errors

type UnathuarizedError struct{}

func NewUnauthorizedError() *UnathuarizedError {
	return &UnathuarizedError{}
}

func (e *UnathuarizedError) Error() string {
	return "Unauthorized"
}
