package method

import "net/http"

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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("login", "*")
	(*w).Header().Set("password", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
