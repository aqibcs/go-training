package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-training/auth"
	"go-training/consts"
	"go-training/db/models/crud"
	"go-training/db/models/object"

	"github.com/go-chi/chi"
)

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == consts.API_KEY {
			token, err := auth.CreateJWt()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}

func GetAllObjects(w http.ResponseWriter, r *http.Request) {
	objects, err := crud.GetAllObjects()
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

	obj, err := crud.GetObjectByID(uint(objectID))
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

	if err := crud.CreateObject(&obj); err != nil {
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

	updatedObject, err := crud.UpdateObject(uint(objectID), &obj)
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

	if err := crud.DeleteObject(uint(objectID)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Return response indicating success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Object deleted successfully"}`))
}
