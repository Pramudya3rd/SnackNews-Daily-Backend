package errors

import "fmt"

// NotFoundError represents an error when a resource is not found.
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

// ValidationError represents an error for validation failures.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}

// InternalServerError represents an internal server error.
type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Internal server error: %s", e.Message)
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(resource, id string) error {
	return &NotFoundError{Resource: resource, ID: id}
}

// NewValidationError creates a new ValidationError.
func NewValidationError(field, message string) error {
	return &ValidationError{Field: field, Message: message}
}

// NewInternalServerError creates a new InternalServerError.
func NewInternalServerError(message string) error {
	return &InternalServerError{Message: message}
}