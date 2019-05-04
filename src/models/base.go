package models

type ModelInterface interface {
	SaveNew()
	Update()
	Delete()
	FindById(int)
}

type Model struct {
	ID uint
}
