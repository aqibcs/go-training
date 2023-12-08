package crud

import (
	"go-training/db"
	models "go-training/models/object"
)

func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	if err := db.Conn().Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func GetEmployeeByID(id uint) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func CreateEmployee(employee *models.Employee) error {
	return db.Conn().Create(employee).Error
}

func UpdateEmployee(id uint, updatedEmployee *models.Employee) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}

	// Update object fields
	employee.Name = updatedEmployee.Name
	employee.Age = updatedEmployee.Age
	employee.City = updatedEmployee.City

	if err := db.Conn().Save(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func DeleteEmployee(id uint) error {
	return db.Conn().Delete(&models.Employee{}, id).Error
}
