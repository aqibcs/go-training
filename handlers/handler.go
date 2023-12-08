package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-training/auth"
	"go-training/consts"
	"go-training/crud"
	models "go-training/models/object"

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

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := crud.GetAllEmployees()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, employees)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(chi.URLParam(r, "object_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	obj, err := crud.GetEmployeeByID(uint(employeeID))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	sendResponse(w, http.StatusOK, obj)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var obj models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := crud.CreateEmployee(&obj); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, obj)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(chi.URLParam(r, "employee_id"))
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

	updatedEmployee, err := crud.UpdateEmployee(uint(employeeID), &obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, updatedEmployee)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(chi.URLParam(r, "employee_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := crud.DeleteEmployee(uint(employeeID)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, "Delete Employee successfully")
}
