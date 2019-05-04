package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simple-to-do/src/api/method"
	"simple-to-do/src/db"
	wsmethod "simple-to-do/src/websocket/method"
	"time"
)

func main() {
	dir := "./front/build"
	router := mux.NewRouter()

	router.PathPrefix("/api/v1/auth").HandlerFunc(method.Login)
	router.PathPrefix("/api/v1/register").HandlerFunc(method.Register)
	router.PathPrefix("/ws/v1/todo").HandlerFunc(wsmethod.TodoHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer db.Close()
	log.Fatal(srv.ListenAndServe())
}
