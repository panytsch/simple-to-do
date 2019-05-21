package main

import (
	"log"
	"net/http"
	"simple-to-do/src/api/method"
	"simple-to-do/src/db"
	"simple-to-do/src/middleware"
	"simple-to-do/src/websocket"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	dir := "./front/build"
	router := mux.NewRouter()

	router.PathPrefix("/api/v1/auth").HandlerFunc(method.Login).Methods("POST", "OPTIONS")
	router.PathPrefix("/api/v1/register").HandlerFunc(method.Register).Methods("POST", "OPTIONS")
	router.PathPrefix("/ws/v1/todo").HandlerFunc(websocket.TodoHandler).Methods("GET", "OPTIONS")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))

	router.Use(middleware.IsThisRouteAllowedMiddleware)
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
