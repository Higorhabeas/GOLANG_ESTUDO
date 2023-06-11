package main

import (
	"log"
	"net/http"

	servidor "crud/24.2-servidor"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", router))

}
