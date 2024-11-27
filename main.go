package main

import (
	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	config.Migrate()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
