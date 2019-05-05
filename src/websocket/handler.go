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
			todos := user.GetAllTodos()
			for _, todo := range todos {
				clientResponse.Todos = append(clientResponse.Todos, todo.GetResponseTodo())
			}
			log.Println("Type ", TypeConnect, " is ok. response: ", clientResponse)
		case TypeAdd:
			todo := models.Todo{}
			todo.BuildFromResponseTodo(clientRequest.Todo)
			todo.User = *user
			todo.UserID = user.ID
			errValidate := todo.Validate()
			if errValidate != nil {
				clientResponse.Fail = true
				clientResponse.Message = errValidate.Error()
				break
			}
			todo.SaveNew()
			clientResponse.Todos = []models.ResponseTodo{todo.GetResponseTodo()}
		}
		log.Println("Type ", clientRequest.Type, "| response: ", clientResponse)
		_ = connection.WriteJSON(clientResponse)
	}
}
