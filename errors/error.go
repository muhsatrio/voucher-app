package customerror

import "net/http"

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func BadRequest(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusBadRequest, Message: msg}
}

func NotFound(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusNotFound, Message: msg}
}

func Internal(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusInternalServerError, Message: msg}
}
