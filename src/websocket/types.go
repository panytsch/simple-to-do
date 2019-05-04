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
		Type   Type
		TodoID uint
		Todo   models.Todo
		Todos  []models.Todo
		Token  string
	}
	ClientResponse struct {
		Todos   []models.Todo
		Fail    bool
		Message string
	}
)

//Type enum
const (
	TypeDelete  Type = "delete"
	TypeUpdate  Type = "update"
	TypeAdd     Type = "add"
	TypeConnect Type = "connect"
)
