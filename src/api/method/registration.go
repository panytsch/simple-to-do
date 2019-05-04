package method

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"simple-to-do/src/api/entities"
	"simple-to-do/src/helpers"
	"simple-to-do/src/models"
)

//register new user or return ErrorInMethodInterface even if found matched user
func registerNewUser(login string, password string) (*models.User, ErrorInMethodInterface) {
	user := models.FindByNamePass(login, password)
	if user != nil {
		err := &ErrorInMethod{}
		err.SetError(errors.New("user already exist"))
		err.SetCode(http.StatusFound)
		return nil, err
	}

	user = &models.User{}
	user.Login = login
	user.Password = password
	user.Token = helpers.RandStringByLength(models.TokenLength)
	err := user.Validate()
	if err != nil {
		err := &ErrorInMethod{}
		err.SetError(err)
		err.SetCode(http.StatusForbidden)
		return nil, err
	}

	user.SaveNew()
	return user, nil
}

func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	user, err := registerNewUser(request.Header.Get("login"), request.Header.Get("password"))
	if err != nil {
		log.Println("register catch err: ", err.Error(), " code: ", err.GetCode())

		response := entities.ErrorResponse{}
		response.Message = err.Error()
		writer.WriteHeader(err.GetCode())
		jsonResponse, _ := json.Marshal(response)
		_, _ = writer.Write(jsonResponse)
		return
	}

	if user == nil {
		log.Println("register logic catch err: user is nil")

		response := entities.ErrorResponse{}
		response.Message = "something went wrong"
		writer.WriteHeader(http.StatusBadGateway)
		jsonResponse, _ := json.Marshal(response)
		_, _ = writer.Write(jsonResponse)
		return
	}

	response := entities.SuccessfulRegister{}
	response.Message = "You are successfully registered"
	response.Token = user.Token
	writer.WriteHeader(http.StatusCreated)
	jsonResponse, _ := json.Marshal(response)
	_, _ = writer.Write(jsonResponse)
	log.Println("new user successfully registered")
}
