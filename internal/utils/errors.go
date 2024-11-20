package utils

// AppError is the base error type for our application
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Error codes
const (
	CodeNotFound      = "NOT_FOUND"
	CodeUnauthorized  = "UNAUTHORIZED"
	CodeForbidden     = "FORBIDDEN"
	CodeBadRequest    = "BAD_REQUEST"
	CodeValidation    = "VALIDATION_ERROR"
	CodeDuplicate     = "DUPLICATE_ENTRY"
	CodeInternalError = "INTERNAL_ERROR"
)

// Error constructors
func NewAppError(code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Code:    CodeNotFound,
		Message: resource + " not found",
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    CodeUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:    CodeForbidden,
		Message: message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    CodeBadRequest,
		Message: message,
	}
}

func NewValidationError(field, message string) *AppError {
	return &AppError{
		Code:    CodeValidation,
		Message: message,
		Field:   field,
	}
}

func NewDuplicateEntryError(resource string) *AppError {
	return &AppError{
		Code:    CodeDuplicate,
		Message: resource + " already exists",
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Code:    CodeInternalError,
		Message: message,
	}
}

// Helper functions
func IsNotFound(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == CodeNotFound
	}
	return false
}

func IsValidation(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == CodeValidation
	}
	return false
}

func IsForbidden(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == CodeForbidden
	}
	return false
}
