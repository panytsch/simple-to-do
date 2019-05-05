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
	setupResponse(&writer, request)
	if request.Method == "OPTIONS" {
		return
	}
	writer.Header().Add("Content-Type", "application/json")

	postRequest := entities.PostRegisterRequest{}
	decoder := json.NewDecoder(request.Body)
	er := decoder.Decode(&postRequest)
	if er != nil {
		log.Println("wrong request: ", er.Error())
		response := entities.ErrorResponse{}
		response.Message = er.Error()
		writer.WriteHeader(http.StatusForbidden)
		jsonResponse, _ := json.Marshal(response)
		_, _ = writer.Write(jsonResponse)
		return
	}

	user, err := registerNewUser(postRequest.Login, postRequest.Password)
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
