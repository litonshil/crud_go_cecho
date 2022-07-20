package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
)

func Registration(c echo.Context) error {
	var user = new(models.User)
	
	//bind
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	if err := repository.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "user created successfullys")
}

func Health(c echo.Context) error {
	
	return c.JSON(http.StatusCreated, "user created successfullys")
}
