package main

import (
	"go-training/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/csv-json-conv/{path}", handlers.CSVtoJSONHandler).Methods("GET")
	r.HandleFunc("/hello/{name}", handlers.HelloHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
