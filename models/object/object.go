package models

import (
	"go-training/db"

	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	Name string
	Age  int
	City string
}

func GetAllObjects() ([]Object, error) {
	var objects []Object
	if err := db.GetDB().Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func GetObjectByID(id uint) (Object, error) {
	var object Object
	if err := db.GetDB().First(&object, id).Error; err != nil {
		return Object{}, err
	}
	return object, nil
}

func CreateObject(object *Object) error {
	if err := db.GetDB().Create(object).Error; err != nil {
		return err
	}
	return nil
}

func UpdateObject(id uint, updatedObject *Object) (Object, error) {
	var object Object
	if err := db.GetDB().First(&object, id).Error; err != nil {
		return Object{}, err
	}

	// Update object fields
	object.Name = updatedObject.Name
	object.Age = updatedObject.Age
	object.City = updatedObject.City

	if err := db.GetDB().Save(&object).Error; err != nil {
		return Object{}, err
	}

	return object, nil
}

func DeleteObject(id uint) error {
	if err := db.GetDB().Delete(&Object{}, id).Error; err != nil {
		return err
	}
	return nil
}