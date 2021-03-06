package method

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"simple-to-do/src/api/entities"
	"simple-to-do/src/models"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	setupResponse(&writer, request)
	if request.Method == "OPTIONS" {
		return
	}
	writer.Header().Add("Content-Type", "application/json")

	postRequest := entities.PostLoginRequest{}
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
	user, err := login(postRequest.Login, postRequest.Password)
	if err != nil {
		log.Println("login catch err: ", err.Error(), " code: ", err.GetCode())

		response := entities.ErrorResponse{}
		response.Message = err.Error()
		writer.WriteHeader(err.GetCode())
		jsonResponse, _ := json.Marshal(response)
		_, _ = writer.Write(jsonResponse)
		return
	}

	if user == nil {
		log.Println("login logic catch err: user is nil")

		response := entities.ErrorResponse{}
		response.Message = "something went wrong"
		writer.WriteHeader(http.StatusBadGateway)
		jsonResponse, _ := json.Marshal(response)
		_, _ = writer.Write(jsonResponse)
		return
	}

	response := entities.SuccessfulLogin{}
	response.Message = "Login successfully"
	response.Token = user.Token
	writer.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(response)
	_, _ = writer.Write(jsonResponse)
	log.Println("successful login")
}

func login(login string, password string) (*models.User, ErrorInMethodInterface) {
	if login == "" {
		err := &ErrorInMethod{}
		err.SetError(errors.New("please send your login"))
		err.SetCode(http.StatusNotFound)
		return nil, err
	}

	if password == "" {
		err := &ErrorInMethod{}
		err.SetError(errors.New("please send your password"))
		err.SetCode(http.StatusNotFound)
		return nil, err
	}

	user := models.FindByNamePass(login, password)
	if user == nil {
		err := &ErrorInMethod{}
		err.SetError(errors.New("user not found. login: " + login + " pass: " + password))
		err.SetCode(http.StatusNotFound)
		return nil, err
	}
	return user, nil
}
