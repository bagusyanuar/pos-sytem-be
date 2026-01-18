package response

import "time"

// APIResponse template generic untuk semua response
type APIResponse[T any] struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Data      T         `json:"data,omitempty"`
	Error     *APIError `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type APIError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func Success[T any](message string, data T) APIResponse[T] {
	return APIResponse[T]{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

func Error(message string, err error) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Request failed",
		Error: &APIError{
			Message: message,
			Details: err.Error(),
		},
		Timestamp: time.Now(),
	}
}

func ErrorWithCode(code, message string, details any) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Request failed",
		Error: &APIError{
			Code:    code,
			Message: message,
			Details: details,
		},
		Timestamp: time.Now(),
	}
}

func ValidationError(details any) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Validation failed",
		Error: &APIError{
			Code:    "VALIDATION_ERROR",
			Message: "Invalid input data",
			Details: details,
		},
		Timestamp: time.Now(),
	}
}

// NotFound membuat response untuk resource not found
func NotFound(resource string) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Resource not found",
		Error: &APIError{
			Code:    "NOT_FOUND",
			Message: resource + " not found",
		},
		Timestamp: time.Now(),
	}
}

func BadRequest(message string) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Bad Request",
		Error: &APIError{
			Code:    "BAD_REQUEST",
			Message: message,
		},
		Timestamp: time.Now(),
	}
}

// Unauthorized membuat response untuk unauthorized access
func Unauthorized(message string) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Unauthorized",
		Error: &APIError{
			Code:    "UNAUTHORIZED",
			Message: message,
		},
		Timestamp: time.Now(),
	}
}

func Forbidden(message string) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Forbidden",
		Error: &APIError{
			Code:    "FORBIDDEN_ACCESS",
			Message: message,
		},
		Timestamp: time.Now(),
	}
}

// InternalServerError membuat response untuk internal server error
func InternalServerError(err error) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Message: "Internal server error",
		Error: &APIError{
			Code:    "INTERNAL_ERROR",
			Message: "An unexpected error occurred",
			Details: err.Error(),
		},
		Timestamp: time.Now(),
	}
}
