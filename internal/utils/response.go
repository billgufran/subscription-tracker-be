package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse creates a success response with optional data
func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// SuccessMessageResponse creates a success response with a message
func SuccessMessageResponse(message string) Response {
	return Response{
		Success: true,
		Message: message,
	}
}

// ErrorResponse creates an error response with a message
func ErrorResponse(message string) Response {
	return Response{
		Success: false,
		Error:   message,
	}
}

// ValidationErrorResponse creates a response for validation errors
func ValidationErrorResponse(errors interface{}) Response {
	return Response{
		Success: false,
		Error:   "Validation failed",
		Data:    errors,
	}
}
