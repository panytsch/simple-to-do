package models

import (
	"errors"
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
		return errors.New("wrong user ID")
	}

	if user.ID == todo.User.ID {
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
