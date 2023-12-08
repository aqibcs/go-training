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

func CreateEmployee(object *models.Employee) error {
	return db.Conn().Create(object).Error
}

func UpdateEmployee(id uint, updatedEmploye *models.Employee) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}

	if err := db.Conn().Save(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func DeleteEmployee(id uint) error {
	return db.Conn().Delete(&models.Employee{}, id).Error
}
