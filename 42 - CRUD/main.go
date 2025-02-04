package main

import (
	"crud/app"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	endpoint := "/usuarios"

	router.HandleFunc(endpoint, app.GetAll).Methods(http.MethodGet)
	router.HandleFunc(endpoint+"/{id}", app.GetById).Methods(http.MethodGet)
	router.HandleFunc(endpoint, app.Insert).Methods(http.MethodPost)
	router.HandleFunc(endpoint+"/{id}", app.Update).Methods(http.MethodPut)
	router.HandleFunc(endpoint+"/{id}", app.Delete).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
