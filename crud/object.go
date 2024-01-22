package crud

import (
	"go-training/db"
	"go-training/models/object"
)

// GetAllEmployees retrieves all employees from the database.
func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	if err := db.Conn().Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// GetEmployeeByID retrieves a single employee by ID from the database.
func GetEmployeeByID(id uint) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

// CreateEmployee creates a new employee in the database.
func CreateEmployee(employee *models.Employee) error {
	return db.Conn().Create(employee).Error
}

// UpdateEmployee updates an existing employee in the database.
func UpdateEmployee(id uint, updatedEmployee *models.Employee) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}

	// Update employee fields
	employee.Name = updatedEmployee.Name
	employee.Age = updatedEmployee.Age
	employee.City = updatedEmployee.City

	if err := db.Conn().Save(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

// DeleteEmployee deletes an employee from the database based on its ID.
func DeleteEmployee(id uint) error {
	return db.Conn().Delete(&models.Employee{}, id).Error
}
