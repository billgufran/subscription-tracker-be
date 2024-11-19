package utils

import "errors"

// Common errors
var (
	ErrNotFound           = errors.New("resource not found")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrBadRequest         = errors.New("bad request")
	ErrInvalidInput       = errors.New("invalid input")
	ErrDuplicateEntry     = errors.New("duplicate entry")
	ErrInternalServer     = errors.New("internal server error")
	ErrValidation         = errors.New("validation error")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// DatabaseError wraps database-related errors
type DatabaseError struct {
	Err error
}

func (e *DatabaseError) Error() string {
	return "database error: " + e.Err.Error()
}

// ValidationError holds validation errors
type ValidationError struct {
	Field   string
	Message string
}

// ValidationErrors is a slice of ValidationError
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return "validation failed"
	}
	return ve[0].Message
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

// IsNotFoundError checks if the error is a not found error
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsDuplicateError checks if the error is a duplicate entry error
func IsDuplicateError(err error) bool {
	return errors.Is(err, ErrDuplicateEntry)
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(ValidationErrors)
	return ok || errors.Is(err, ErrValidation)
}
