package apperror

// AppError cusom defined app error
type AppError struct {
	Code    uint64
	Cause   string
	Message string
}

// New returns new error object
func New(code uint64, cause string, message string) *AppError {
	return &AppError{
		Code:    code,
		Cause:   cause,
		Message: message,
	}
}
