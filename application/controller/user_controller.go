package controller

import (
	"backend/core/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface {
	GetUser(c echo.Context) error
}

type userController struct {
	userService service.UserService
}

func UserControllerConstructor(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) GetUser(c echo.Context) error {
	userId := c.Param("userId")
	user, err := controller.userService.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
