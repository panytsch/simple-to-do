package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simple-to-do/src/api/method"
	"simple-to-do/src/db"
	"time"
)

func main() {
	dir := "./front/build"
	router := mux.NewRouter()

	router.PathPrefix("/api/auth").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		method.Login(writer, request)
	})
	router.PathPrefix("/api/register").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		method.Register(writer, request)
	})

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
