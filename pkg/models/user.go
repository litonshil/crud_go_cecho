package models
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       	int			`gorm:"primary_key;AUTO_INCREMENT"`
	Name     	string      `json:"name"`
	Address		string		`json:"address"`	
}