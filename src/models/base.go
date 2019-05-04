package models

import "github.com/jinzhu/gorm"

type ModelInterface interface {
	SaveNew()
	Update()
	Delete()
	FindById(int)
}

type Model struct {
	gorm.Model
}
