package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + name + "!"))
}
