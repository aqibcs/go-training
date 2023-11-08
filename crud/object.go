package crud

import (
	"go-training/db"
	models "go-training/models/object"
)

func GetAllObjects() ([]models.Employee, error) {
	var objects []models.Employee
	if err := db.Conn().Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func GetObjectByID(id uint) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func CreateObject(object *models.Employee) error {
	return db.Conn().Create(object).Error
}

func UpdateObject(id uint, updatedObject *models.Employee) (models.Employee, error) {
	var employee models.Employee
	if err := db.Conn().First(&employee, id).Error; err != nil {
		return models.Employee{}, err
	}

	// Update object fields
	employee.Name = updatedObject.Name
	employee.Age = updatedObject.Age
	employee.City = updatedObject.City

	if err := db.Conn().Save(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func DeleteObject(id uint) error {
	return db.Conn().Delete(&models.Employee{}, id).Error
}
