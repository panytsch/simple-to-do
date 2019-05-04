package entities

type Response struct {
	Message string
}

type ErrorResponse struct {
	Response
}

type SuccessfulRegister struct {
	Response
	Token string
}

type SuccessfulLogin struct {
	Response
	Token string
}
