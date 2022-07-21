package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
)

func User(e *echo.Echo) {
	sub := e.Group("/user")
	sub.POST("/registration", controllers.Registration)
	sub.POST("/login", controllers.Login)
	sub.GET("/users", controllers.GetAllUsers)
	sub.GET("/search/:id", controllers.GetAUsers)
	sub.PUT("/update/:id", controllers.UpdateUser)
	sub.DELETE("/delete/:id", controllers.DeleteUser)

}
