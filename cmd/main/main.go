package main

import (
	"backend/application/config"
	"backend/application/route"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConfEnv()

	route.Router(e)

	e.Logger.Fatal(e.Start(":" + config.GetPort()))
}
