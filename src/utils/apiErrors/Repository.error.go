package apiErrors

var _ IApiErrors = (*RepositoryError)(nil)

type RepositoryError struct {
	Message string
}

func (e *RepositoryError) Error() string {
	return e.Message
}

func NewRepositoryError(message string) *RepositoryError {
	return &RepositoryError{Message: message}
}
