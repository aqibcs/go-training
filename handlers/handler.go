package handlers

import (
	"go-training/auth"
	"go-training/consts"
	"go-training/crud"
	"go-training/models/object"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetJwt handles the HTTP GET request for retrieving a JWT.
func GetJwt(c echo.Context) error {
	apiKey := c.Request().Header.Get("Access")
	if apiKey != "" && apiKey == consts.APIKey {
		token, err := auth.CreateJWT()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.String(http.StatusOK, token)
	}

	return c.String(http.StatusUnauthorized, "Unauthorized")
}

// sendResponse sends an HTTP response with the specified status code and data in JSON format.
func sendResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}

// GetAllEmployees handles the HTTP GET request for retrieving all employees.
func GetAllEmployees(c echo.Context) error {
	employees, err := crud.GetAllEmployees()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return sendResponse(c, http.StatusOK, employees)
}

// GetEmployeeByID handles the HTTP GET request for retrieving a specific employee by ID.
func GetEmployeeByID(c echo.Context) error {
	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	employee, err := crud.GetEmployeeByID(uint(employeeID))
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return sendResponse(c, http.StatusOK, employee)
}

// CreateEmployee handles the HTTP POST request for creating a new employee.
func CreateEmployee(c echo.Context) error {
	var employee models.Employee
	if err := c.Bind(&employee); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if err := crud.CreateEmployee(&employee); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return sendResponse(c, http.StatusCreated, employee)
}

// UpdateEmployee handles the HTTP PATCH request for updating an existing employee.
func UpdateEmployee(c echo.Context) error {
	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	var employee models.Employee
	if err := c.Bind(&employee); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	updatedEmployee, err := crud.UpdateEmployee(uint(employeeID), &employee)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return sendResponse(c, http.StatusOK, updatedEmployee)
}

// DeleteEmployee handles the HTTP DELETE request for deleting an employee by its ID.
func DeleteEmployee(c echo.Context) error {
	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if err := crud.DeleteEmployee(uint(employeeID)); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.NoContent(http.StatusNoContent)
}
