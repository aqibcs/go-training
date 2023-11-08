package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-training/crud"
	models "go-training/models/object"

	"github.com/go-chi/chi"
)

func sendResponse(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func GetAllObjects(w http.ResponseWriter, r *http.Request) {
	objects, err := crud.GetAllObjects()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, objects)
}

func GetObjectByID(w http.ResponseWriter, r *http.Request) {
	objectID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	obj, err := crud.GetObjectByID(uint(objectID))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	sendResponse(w, http.StatusOK, obj)
}

func CreateObject(w http.ResponseWriter, r *http.Request) {
	var obj models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := crud.CreateObject(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, obj)
}

func UpdateObject(w http.ResponseWriter, r *http.Request) {
	objectID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var obj models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updatedObject, err := crud.UpdateObject(uint(objectID), &obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, updatedObject)
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	objectID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := crud.DeleteObject(uint(objectID)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, "Delete object successfully")
}
