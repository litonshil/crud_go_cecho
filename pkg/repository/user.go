package repository

import (
	"fmt"

	"github.com/litonshil/crud_go_echo/pkg/database"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

var db = database.GetDB()

func CreateUser(user *models.User) error {
	err := db.Create(&user).Error
	return err
}
func GetAllUsers() ([]models.User, error) {
	var all_users []models.User
	err := db.Find(&all_users).Error
	fmt.Println(all_users)
	return all_users, err

}

func GetAUsers(id int) ([]models.User, error) {
	var user []models.User
	err := db.Where("id = ?", id).Find(&user).Error
	fmt.Println("user",user)
	return user, err
}
