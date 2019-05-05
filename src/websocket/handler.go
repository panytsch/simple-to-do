package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"simple-to-do/src/models"
)

var upgrader websocket.Upgrader

func init() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func TodoHandler(writer http.ResponseWriter, request *http.Request) {
	token := request.URL.Query().Get("token")
	if token == "" {
		log.Println("no token")
		return
	}
	user := models.GetUserByToken(token)
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("can't upgrade")
		return
	}

	go readAndWrite(user, connection)
}

func readAndWrite(user *models.User, connection *websocket.Conn) {
	for {
		clientRequest := &ClientRequest{}
		clientResponse := &ClientResponse{}
		err := connection.ReadJSON(clientRequest)
		if err != nil {
			log.Println("can't read json from request: ", clientRequest)
			clientResponse.Message = "can't read json"
			clientResponse.Fail = true
			err := connection.WriteJSON(clientResponse)
			if err != nil {
				log.Println("try to close connection")
				_ = connection.Close()
				return
			}
			continue
		}

		switch clientRequest.Type {
		case TypeConnect:
			typeConnect(clientResponse, user)
		case TypeAdd:
			typeAdd(clientRequest, clientResponse, user)
		case TypeUpdate:
			typeUpdate(clientRequest, clientResponse, user)
		case TypeDelete:
			typeDelete(clientRequest, clientResponse, user)
		}
		log.Println("Type ", clientRequest.Type, "| response: ", clientResponse)
		_ = connection.WriteJSON(clientResponse)
	}
}

func typeDelete(request *ClientRequest, response *ClientResponse, user *models.User) {
	todo := models.Todo{}
	if request.Todo.ID == 0 {
		response.Fail = true
		response.Message = "send todo id"
		log.Println("todo id wasn't sent")
		return
	}
	todo.FindByIdAndUserId(request.Todo.ID, user.ID)
	if todo.ID == 0 {
		response.Fail = true
		response.Message = "todo not found"
		log.Println("todo not found")
		return
	}
	todo.Delete()
	response.Message = "successfully"
	log.Println("todo deleted", todo)
}

func typeUpdate(request *ClientRequest, response *ClientResponse, user *models.User) {
	todo := models.Todo{}
	if request.Todo.ID == 0 {
		response.Fail = true
		response.Message = "send todo id"
		log.Println("todo id wasn't sent")
		return
	}
	todo.FindByIdAndUserId(request.Todo.ID, user.ID)
	if todo.ID == 0 {
		response.Fail = true
		response.Message = "todo not found"
		log.Println("todo not found")
		return
	}
	todo.BuildFromResponseTodo(request.Todo)
	todo.Update()
	response.Todos = []models.ResponseTodo{todo.GetResponseTodo()}
	log.Println("todo updated")
}

func typeAdd(request *ClientRequest, response *ClientResponse, user *models.User) {
	todo := models.Todo{}
	todo.BuildFromResponseTodo(request.Todo)
	todo.User = *user
	todo.UserID = user.ID
	errValidate := todo.Validate()
	if errValidate != nil {
		response.Fail = true
		response.Message = errValidate.Error()
		return
	}
	todo.SaveNew()
	response.Todos = []models.ResponseTodo{todo.GetResponseTodo()}
}

func typeConnect(response *ClientResponse, user *models.User) {
	todos := user.GetAllTodos()
	for _, todo := range todos {
		response.Todos = append(response.Todos, todo.GetResponseTodo())
	}
}
