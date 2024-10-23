package apiErrors

var _ IApiErrors = (*RepositoryError)(nil)

type InvalidPropsError struct {
	Message string
}

func (e *InvalidPropsError) Error() string {
	return e.Message
}

func NewInvalidPropsError(message string) *InvalidPropsError {
	return &InvalidPropsError{Message: message}
}
