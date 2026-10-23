package errors

import "net/http"

type Code string

const (
	// Generic
	CodeInternal        Code = "INTERNAL_SERVER_ERROR"
	CodeBadRequest      Code = "BAD_REQUEST"
	CodeValidation      Code = "VALIDATION_FAILED"
	CodeUnauthorized    Code = "UNAUTHORIZED"
	CodeForbidden       Code = "FORBIDDEN"
	CodeNotFound        Code = "NOT_FOUND"
	CodeConflict        Code = "CONFLICT"
	CodeTooManyRequests Code = "TOO_MANY_REQUESTS"
	CodeUnavailable     Code = "SERVICE_UNAVAILABLE"

	// Auth / user
	CodeInvalidCredentials Code = "INVALID_CREDENTIALS"
	CodeTokenExpired       Code = "TOKEN_EXPIRED"
	CodeEmailTaken         Code = "EMAIL_ALREADY_EXISTS"

	// Wallet / money
	CodeInsufficientFunds Code = "INSUFFICIENT_FUNDS"
	CodeInvalidAmount     Code = "INVALID_AMOUNT"
	CodeAccountFrozen     Code = "ACCOUNT_FROZEN"
	CodeDuplicateRequest  Code = "DUPLICATE_REQUEST" // idempotency key reused
)

type AppError struct {
	StatusCode int    `json:"_"`
	Code       Code   `json:"code"`
	Message    string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppErrro(statusCode int, code Code, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

var (
	ErrInternal     = NewAppErrro(http.StatusInternalServerError, CodeInternal, "Something went wrong, please try again later.")
	ErrBadRequest   = NewAppErrro(http.StatusBadRequest, CodeBadRequest, "Bad request, please check the request payload.")
	ErrValidation   = NewAppErrro(http.StatusUnprocessableEntity, CodeValidation, "Validation failed.")
	ErrUnauthorized = NewAppErrro(http.StatusUnauthorized, CodeUnauthorized, "Unauthorized, please check the credentials.")
	ErrForbidden    = NewAppErrro(http.StatusForbidden, CodeForbidden, "You don't have permission to do this.")
	ErrNotFound     = NewAppErrro(http.StatusNotFound, CodeNotFound, "Resource not found.")
	ErrConflict     = NewAppErrro(http.StatusConflict, CodeConflict, "Resource already exists.")
)
