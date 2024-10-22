package apiErrors

var _ IApiErrors = (*InvalidPropsErros)(nil)

type InvalidPropsErros struct {
	Message string
}

func (e *InvalidPropsErros) Error() string {
	return e.Message
}

func NewInvalidPropsException(message string) *InvalidPropsErros {
	return &InvalidPropsErros{Message: message}
}
