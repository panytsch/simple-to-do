package models

import (
	"errors"
	"simple-to-do/src/db"
	"simple-to-do/src/helpers"
)

const (
	NameLength     = 30
	PasswordLength = 100
	TokenLength    = 300
)

func init() {
	db.GetDB().AutoMigrate(&User{})
}

type User struct {
	Model
	Login    string `gorm:"type:varchar(30);unique_index;not null"`  //NAME_LENGTH
	Password string `gorm:"type:varchar(100); not null"`             //PASSWORD_LENGTH
	Token    string `gorm:"type:varchar(300);unique_index;not null"` //TOKEN_LENGTH
	Todos    []Todo
}

type ResponseUser struct {
	Login string
	Token string
}

func (User) TableName() string {
	return "users"
}

func FindByNamePass(name string, pass string) *User {
	var user = &User{}
	db.GetDB().Where("login = ? and password = ? and deleted_at is null", name, pass).Find(user)
	if user.ID != 0 {
		return user
	}
	return nil
}

func (u *User) Validate() error {
	if len(u.Login) > NameLength {
		return errors.New("login is too large")
	}

	if len(u.Password) > PasswordLength {
		return errors.New("password is too large")
	}

	if len(u.Token) > TokenLength {
		return errors.New("token is too large")
	}

	if !CheckToken(u.Token) {
		u.Token = helpers.RandStringByLength(TokenLength)
		return u.Validate() //be careful dude :)
	}

	if !CheckLogin(u.Login) {
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

func (u *User) FindById(id uint) {
	db.GetDB().Where("id = ? and deleted_at is null", id).Find(u)
}

func GetUserByToken(token string) *User {
	user := &User{}
	db.GetDB().Where("token = ? and deleted_at is null", token).Find(user)
	return user
}

func (u *User) GetAllTodos() []Todo {
	var todos []Todo
	selectFields := []string{"id", "created_at", "updated_at", "title", "description", "is_done"}
	db.GetDB().Where("user_id = ? and deleted_at is null", u.ID).Select(selectFields).Find(&todos)
	return todos
}

func (u *User) GetResponseUser() ResponseUser {
	user := ResponseUser{}
	user.Login = u.Login
	user.Token = u.Token
	return user
}
