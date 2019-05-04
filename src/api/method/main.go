package method

type ErrorInMethodInterface interface {
	error
	GetCode() int
	SetCode(int)
	SetError(error)
}

type ErrorInMethod struct {
	error
	code int
}

func (ErrorInMethod *ErrorInMethod) GetCode() int {
	return ErrorInMethod.code
}

func (ErrorInMethod *ErrorInMethod) SetCode(code int) {
	ErrorInMethod.code = code
}

func (ErrorInMethod *ErrorInMethod) SetError(error error) {
	ErrorInMethod.error = error
}
