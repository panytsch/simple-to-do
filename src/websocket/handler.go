package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"simple-to-do/src/models"
)

var todoMapClient map[string]*Client
var upgrader websocket.Upgrader

func init() {
	todoMapClient = make(map[string]*Client)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func TodoHandler(writer http.ResponseWriter, request *http.Request) {
	//token := request.Header.Get("token")
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

	for {
		clientRequest := &ClientRequest{}
		err := connection.ReadJSON(clientRequest)
		if err != nil {
			log.Println("can't read json")
			return
		}

		clientResponse := &ClientResponse{}

		switch clientRequest.Type {
		case TypeConnect:
			clientResponse.Todos = user.GetAllTodos()
			_ = connection.WriteJSON(clientResponse)
		}
	}
}
