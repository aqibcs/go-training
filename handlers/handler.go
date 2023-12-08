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

func GetJwt(c echo.Context) error {
	apiKey := c.Request().Header.Get("Access")
	if apiKey != "" && apiKey == consts.API_KEY {
		token, err := auth.CreateJWt()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, token)
	}

	return c.String(http.StatusUnauthorized, "Unauthorized")
}

// GetAllEmployees returns all employees.
func GetAllEmployees(c echo.Context) error {
	employees, err := crud.GetAllEmployees()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, employees)
}

// GetEmployeeByID returns an employee by ID.
func GetEmployeeByID(c echo.Context) error {
	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid employee ID")
	}

	employee, err := crud.GetEmployeeByID(uint(employeeID))
	if err != nil {
		return c.String(http.StatusNotFound, "Employee not found")
	}

	return c.JSON(http.StatusOK, employee)
}

// CreateEmployee adds a new employee.
func CreateEmployee(c echo.Context) error {
	var newEmployee models.Employee
	if err := c.Bind(&newEmployee); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	if err := crud.CreateEmployee(&newEmployee); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusCreated, newEmployee)
}

// UpdateEmployee updates an employee by ID.
func UpdateEmployee(c echo.Context) error {
	employeeId, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid employee ID")
	}

	var updatedEmployee models.Employee
	if err := c.Bind(&updatedEmployee); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	employee, err := crud.UpdateEmployee(uint(employeeId), &updatedEmployee)
	if err != nil {
		return c.String(http.StatusNotFound, "Employee not found")
	}

	return c.JSON(http.StatusOK, employee)
}

// DeleteEmployee deletes an employee by ID.
func DeleteEmployee(c echo.Context) error {
	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid employee ID")
	}

	if err := crud.DeleteEmployee(uint(employeeID)); err != nil {
		return c.String(http.StatusNotFound, "Employee not found")
	}

	return echo.NewHTTPError(http.StatusOK, "Employee delete successfully")
}
