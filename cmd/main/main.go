package main

import (
	"backend/application/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":"))
}
