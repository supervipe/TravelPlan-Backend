package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}
