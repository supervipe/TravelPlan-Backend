package controller

import (
	"backend/domain/dto"
	"backend/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface {
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUserPassword(c echo.Context) error
	DeleteUser(c echo.Context) error
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *userController) CreateUser(c echo.Context) error {
	user := new(dto.CreateUserDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	userDTO, err := controller.userService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, userDTO)
}

func (controller *userController) UpdateUserPassword(c echo.Context) error {
	userId := c.Param("userId")
	user := new(dto.UpdateUserPasswordDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	userDTO, err := controller.userService.UpdateUserPassword(user, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, userDTO)
}

func (controller *userController) DeleteUser(c echo.Context) error {
	userId := c.Param("userId")
	err := controller.userService.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted successfully"})
}
