package route

import (
	"backend/api/controller"
	"backend/infra/repository"
	"backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, db *gorm.DB) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepository := repository.UserRepositoryConstructor(db)
	userService := service.UserServiceConstructor(userRepository)
	userController := controller.UserControllerConstructor(userService)
	e.GET("/users/:userId", userController.GetUser)
	e.POST("/users", userController.CreateUser)
	e.PUT("/users/:userId", userController.UpdateUserPassword)
	e.DELETE("/users/:userId", userController.DeleteUser)
}
