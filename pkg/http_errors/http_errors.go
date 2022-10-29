package http_errors

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrWrongCredentials      = errors.New("wrong credentials")
	ErrNotFound              = errors.New("not found")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrForbidden             = errors.New("forbidden")
	ErrPermissionDenied      = errors.New("permission denied")
	ErrExpiredCSRFError      = errors.New("expired csrf token")
	ErrWrongCSRFToken        = errors.New("wrong csrf token")
	ErrCSRFNotPresented      = errors.New("csrf not presented")
	ErrNotRequiredFields     = errors.New("no such required fields")
	ErrBadQueryParams        = errors.New("invalid query params")
	ErrInternalServerError   = errors.New("internal server error")
	ErrRequestTimeoutError   = errors.New("request timeout")
	ErrExistsEmailError      = errors.New("user with given email already exists")
	ErrInvalidJWTToken       = errors.New("invalid jwt token")
	ErrInvalidJWTClaims      = errors.New("invalid jwt claims")
	ErrNotAllowedImageHeader = errors.New("not allowed image header")
	ErrNoCookie              = errors.New("not found cookie header")
)

// Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Rest error struct
type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

// Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

// New Rest Error
func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

// New Rest Error With Message
func NewRestErrorWithMessage(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

// New Rest Error From Bytes
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// New Bad Request Error
func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  ErrBadRequest.Error(),
		ErrCauses: causes,
	}
}

// New Not Found Error
func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  ErrNotFound.Error(),
		ErrCauses: causes,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  ErrUnauthorized.Error(),
		ErrCauses: causes,
	}
}

// New Forbidden Error
func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  ErrForbidden.Error(),
		ErrCauses: causes,
	}
}

// New Internal Server Error
func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  ErrInternalServerError.Error(),
		ErrCauses: causes,
	}
	return result
}

// Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, ErrNotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeoutError.Error(), err)
	case strings.Contains(err.Error(), "SQLSTATE"):
		return parseSqlErrors(err)
	case strings.Contains(err.Error(), "Field validation"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewRestError(http.StatusBadRequest, ErrBadRequest.Error(), err)
	case strings.Contains(err.Error(), "UUID"):
		return NewRestError(http.StatusBadRequest, err.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
		return NewRestError(http.StatusBadRequest, ErrBadRequest.Error(), err)
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func parseSqlErrors(err error) RestErr {
	if strings.Contains(err.Error(), "23505") {
		return NewRestError(http.StatusBadRequest, ErrExistsEmailError.Error(), err)
	}

	return NewRestError(http.StatusBadRequest, err.Error(), err)
}

func parseValidatorError(err error) RestErr {
	if strings.Contains(err.Error(), "password") {
		return NewRestError(http.StatusBadRequest, "Invalid password, min length 6", err)
	}

	if strings.Contains(err.Error(), "email") {
		return NewRestError(http.StatusBadRequest, "Invalid email", err)
	}

	if strings.Contains(err.Error(), "name") {
		return NewRestError(http.StatusBadRequest, "Invalid name", err)
	}

	return NewRestError(http.StatusBadRequest, err.Error(), err)
}

// Error response
func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
