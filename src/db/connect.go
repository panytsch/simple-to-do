package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("mysql", "root:Test-5949@/chat-go?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatal(err.Error())
	}
	d.Exec("ALTER DATABASE `chat-go` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;")
	db = d
}

func Close() {
	_ = db.Close()
}

func GetDB() *gorm.DB {
	return db
}
