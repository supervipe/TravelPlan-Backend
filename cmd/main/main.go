package main

import (
	"backend/api/config"
	"backend/api/route"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConfEnv()
	db := config.NewDB()

	e := echo.New()
	route.Router(e, db)

	e.Logger.Fatal(e.Start(":" + config.GetPort()))
}
