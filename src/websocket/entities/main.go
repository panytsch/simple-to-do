package entities

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrader websocket.Upgrader

func init() {
	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}
