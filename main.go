package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// session.SetMode(mgo.Monotonic, true)
	// defer session.Close()
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/todo", AddToDo).Methods("POST", "PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
