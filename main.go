package main

import (
	"log"

	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/models"
	"github.com/FelipeCesarSS/estoque-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Produto{})
	if err != nil {
		log.Fatal("Falha ao migrar o banco de dados:", err)
	}

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
