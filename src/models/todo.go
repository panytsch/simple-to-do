package models

import (
	"errors"
	"log"
	"simple-to-do/src/db"
)

const (
	TitleLength       = 50
	DescriptionLength = 500
)

func init() {
	db.GetDB().AutoMigrate(&Todo{})
}

type Todo struct {
	Model
	Title       string `gorm:"type:varchar(50);not null"`
	Description string `gorm:"type:varchar(500);null"`
	IsDone      bool   `gorm:"type:tinyint(1);not null"`
	User        User
	UserID      uint
}

type ResponseTodo struct {
	ID          uint
	Title       string
	Description string
	IsDone      bool
}

func (todo *Todo) Validate() error {
	if len(todo.Title) > TitleLength {
		return errors.New("title is too large")
	}

	if len(todo.Description) > DescriptionLength {
		return errors.New("description is too large")
	}

	user := User{}
	user.FindById(todo.UserID)
	if user.ID == 0 {
		log.Println("created user:", user, "todo:", todo)
		return errors.New("wrong user ID")
	}

	return nil
}

func (todo *Todo) SaveNew() {
	db.GetDB().Create(todo)
}

func (todo *Todo) Update() {
	db.GetDB().Model(todo).Updates(todo)
}

func (todo *Todo) Delete() {
	db.GetDB().Delete(todo)
}

func (todo *Todo) FindById(id int) {
	db.GetDB().Where("id = ?", id).Find(todo)
}

//convert responseTodo to to do
func (todo *Todo) BuildFromResponseTodo(responseTodo ResponseTodo) {
	todo.Description = responseTodo.Description
	todo.Title = responseTodo.Title
	todo.ID = responseTodo.ID
	todo.IsDone = responseTodo.IsDone
}

//Get to do in type response
func (todo *Todo) GetResponseTodo() ResponseTodo {
	responseTodo := ResponseTodo{}
	responseTodo.ID = todo.ID
	responseTodo.IsDone = todo.IsDone
	responseTodo.Title = todo.Title
	responseTodo.Description = todo.Description
	return responseTodo
}
