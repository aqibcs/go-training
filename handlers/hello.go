package handlers

import (
	"encoding/json"
	"go-training/models/request"
	"go-training/models/response"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody request.RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Inavlid requestbody", http.StatusBadRequest)
	}

	response := response.ResponseBody{
		Code:      200,
		Message:   "Welcome " + requestBody.Name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	jsonResponce, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponce)
}
