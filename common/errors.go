package common

// ValidationError is the error returned when an item is invalid with a message
type ValidationError struct {
	Message string
}

// Error returns the error message
func (e *ValidationError) Error() string {
	return e.Message
}
