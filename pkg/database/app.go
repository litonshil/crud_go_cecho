package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_crud?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}

	return db
}