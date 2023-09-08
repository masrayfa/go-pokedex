package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(msg string) NotFoundError {
	return NotFoundError{Error: msg}
}
