package repository

import (
	"github.com/litonshil/crud_go_echo/pkg/database"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

var db = database.GetDB()

func CreateUser(company *models.User) error {
	err := db.Create(&company).Error
	return err
}