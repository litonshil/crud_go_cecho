package controllers

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
)

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
