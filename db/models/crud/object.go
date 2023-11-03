package crud

import (
	"go-training/db"
	"go-training/db/models/object"
)

func GetAllObjects() ([]models.Object, error) {
	var objects []models.Object
	if err := db.GetDB().Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func GetObjectByID(id uint) (models.Object, error) {
	var object models.Object
	if err := db.GetDB().First(&object, id).Error; err != nil {
		return models.Object{}, err
	}
	return object, nil
}

func CreateObject(object *models.Object) error {
	if err := db.GetDB().Create(object).Error; err != nil {
		return err
	}
	return nil
}

func UpdateObject(id uint, updatedObject *models.Object) (models.Object, error) {
	var object models.Object
	if err := db.GetDB().First(&object, id).Error; err != nil {
		return models.Object{}, err
	}

	// Update object fields
	object.Name = updatedObject.Name
	object.Age = updatedObject.Age
	object.City = updatedObject.City

	if err := db.GetDB().Save(&object).Error; err != nil {
		return models.Object{}, err
	}

	return object, nil
}

func DeleteObject(id uint) error {
	if err := db.GetDB().Delete(&models.Object{}, id).Error; err != nil {
		return err
	}
	return nil
}
