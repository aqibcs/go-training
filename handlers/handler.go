package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go-training/models/object"
)

func GetAllObjects(w http.ResponseWriter, r *http.Request) {
	objects, err := models.GetAllObjects()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Serialize objects to JSON
	jsonResponse, err := json.Marshal(objects)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetObjectByID(w http.ResponseWriter, r *http.Request) {
	objectID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	obj, err := models.GetObjectByID(uint(objectID))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Serialize object to JSON
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateObject(w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := models.CreateObject(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Return response indicating success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(obj)
}

func UpdateObject(w http.ResponseWriter, r *http.Request) {
	objectID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var obj models.Object
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updatedObject, err := models.UpdateObject(uint(objectID), &obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Return updated object as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedObject)
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	objectID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := models.DeleteObject(uint(objectID)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Return response indicating success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Object deleted successfully"}`))
}