package handlers

import (
	"net/http"
	"time"
)

const (
	ErrorInvalidJSONBody                  string = "invalid json body"
	ErrorParsingQuery                     string = "error parsing query"
	ErrorAccountNotFound                  string = "account not found"
	ErrorItemNotFound                     string = "item not found"
	ErrorInvalidToken                     string = "invalid token"
	ErrorInvalidUsernamePassword          string = "invalid username or password"
	ErrorProcessingRequest                string = "error processing request"
	ErrorPhoneExists                      string = "phone number exists"
	ErrorExpiredToken                     string = "token has expired"
	ErrorTokenInvalid                     string = "token is invalid"
	ErrorInvalidAuthorizationHeaderFormat string = "invalid authorization header format"
)

type CustomErr struct {
	TimeStamp time.Time `json:"time_stamp"`
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
}

func NewBadRequestError(message string) *CustomErr {
	return &CustomErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusBadRequest,
		Error:     "bad_request",
	}
}

func NewBadExpiredError(message string) *CustomErr {
	return &CustomErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusUnauthorized,
		Error:     "expired_token",
	}
}
func NewInValidTokenError(message string) *CustomErr {
	return &CustomErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusUnauthorized,
		Error:     "invalid_token",
	}
}
func NewInternalServerError(message string) *CustomErr {
	return &CustomErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusInternalServerError,
		Error:     "internal_server_error",
	}
}
func NewNotFoundError(message string) *CustomErr {
	return &CustomErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusNotFound,
		Error:     "not_found",
	}
}
