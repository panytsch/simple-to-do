package models

import (
	"errors"
	"simple-to-do/src/db"
	"simple-to-do/src/helpers"
)

const NAME_LENGTH = 30
const PASSWORD_LENGTH = 100
const TOKEN_LENGTH = 300

func init() {
	db.GetDB().AutoMigrate(&User{})
}

type User struct {
	Model
	Login    string `gorm:"type:varchar(30);unique_index;not null"`  //NAME_LENGTH
	Password string `gorm:"type:varchar(100); not null"`             //PASSWORD_LENGTH
	Token    string `gorm:"type:varchar(300);unique_index;not null"` //TOKEN_LENGTH
}

func (User) TableName() string {
	return "users"
}

func FindByNamePass(name string, pass string) *User {
	var user = &User{}
	db.GetDB().Where("login = ? and password = ?", name, pass).Find(user)
	if user.ID != 0 {
		return user
	}
	return nil
}

func (user *User) Validate() error {
	if len(user.Login) > NAME_LENGTH {
		return errors.New("login is too large")
	}

	if len(user.Password) > PASSWORD_LENGTH {
		return errors.New("password is too large")
	}

	if len(user.Token) > TOKEN_LENGTH {
		return errors.New("token is too large")
	}

	if !CheckToken(user.Token) {
		user.Token = helpers.RandStringByLength(TOKEN_LENGTH)
		return user.Validate() //be careful dude :)
	}

	if !CheckLogin(user.Login) {
		return errors.New("user already exist")
	}

	return nil
}

func CheckToken(token string) bool {
	user := User{}
	db.GetDB().Where("token = ?", token).First(&user)
	return user.ID == 0
}

func CheckLogin(token string) bool {
	user := User{}
	db.GetDB().Where("login = ?", token).First(&user)
	return user.ID == 0
}

func (u *User) SaveNew() {
	db.GetDB().Create(u)
}

func (u *User) Update() {
	db.GetDB().Model(u).Updates(u)
}

func (u *User) Delete() {
	db.GetDB().Delete(u)
}

func (u *User) FindById(id int) {
	db.GetDB().Where("id = ?", id).Find(u)
}