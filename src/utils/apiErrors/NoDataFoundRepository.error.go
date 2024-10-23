package apiErrors

var _ IApiErrors = (*RepositoryError)(nil)

type NoDataFoundRepositoryError struct {
	Message string
}

func (e *NoDataFoundRepositoryError) Error() string {
	return e.Message
}

func NewNoDataFoundRepositoryError(message string) *NoDataFoundRepositoryError {
	return &NoDataFoundRepositoryError{Message: message}
}
