package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/database"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
)

var db = database.GetDB()

func Registration(c echo.Context) error {
	var user = new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := repository.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "user created successfullys")
}

func GetAllUsers(c echo.Context) error {

	res, err := repository.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func GetAUsers(c echo.Context) error {
	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	res, err := repository.GetAUsers(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func checkEmptyUserField(user *models.User,old_user *models.User) *models.User {
	if (user.Name == "") {
		user.Name = old_user.Name
	}
	if user.Address == "" {
		user.Address = old_user.Address
	}
	if user.Email == "" {
		user.Email = old_user.Email
	}
	if user.Type == "" {
		user.Type = old_user.Type
	}
	if user.Password == "" {
		user.Password = old_user.Password
	}
	return user
}

func UpdateUser(c echo.Context) error {

	var user = new(models.User)
	var old_user = new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}


	id := c.Param("id")

	user_id, _ := strconv.Atoi(id)

	old_err := db.Model(old_user).Where("id = ?", id).Find(&old_user).Error

	if old_err != nil {
		return c.JSON(http.StatusInternalServerError, old_err.Error())
	}

	user.Id = user_id

	checkedUser := checkEmptyUserField(user,old_user)

	res, err := repository.UpdateUser(user_id, checkedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	err := repository.DeleteUser(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "user deleted successfully")
}