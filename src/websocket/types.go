package websocket

import (
	"github.com/gorilla/websocket"
	"simple-to-do/src/models"
)

type (
	Connections []*websocket.Conn
	Client      struct {
		Todos []models.Todo
		Connections
	}
	Type          string
	ClientRequest struct {
		Type  Type
		Todo  models.ResponseTodo
		Todos []models.ResponseTodo
		Token string
	}
	ClientResponse struct {
		Todos   []models.ResponseTodo
		Fail    bool
		Message string
		Type    Type
	}
)

//Type enum
const (
	TypeDelete  Type = "delete"
	TypeUpdate  Type = "update"
	TypeAdd     Type = "add"
	TypeConnect Type = "connect"
)
