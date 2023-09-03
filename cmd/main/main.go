package main

import (
	"backend/application/config"
	"backend/application/route"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConfEnv()

	e := echo.New()
	route.Router(e)

	e.Logger.Fatal(e.Start(":" + config.GetPort()))
}
