package route

import (
	"backend/application/controller"
	"backend/core/service"
	"backend/infra/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepository := repository.UserRepositoryConstructor()
	userService := service.UserServiceConstructor(userRepository)
	userController := controller.UserControllerConstructor(userService)
	e.GET("/users/:userId", userController.GetUser)
	e.POST("/users", userController.CreateUser)
}
