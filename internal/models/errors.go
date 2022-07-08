package models

const (
	InternalLinksServerError = "transaction_server_error"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type ServerError struct {
	Code    string
	Message string
}

func NewLinksError(err error) *ServerError {
	return &ServerError{
		Code:    InternalLinksServerError,
		Message: err.Error(),
	}
}
