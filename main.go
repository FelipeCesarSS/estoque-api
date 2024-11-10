package main

import (
	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Conectar ao banco de dados
	config.ConnectDatabase()

	// Iniciar Echo
	e := echo.New()

	// Configurar rotas
	routes.InitRoutes(e)

	// Iniciar servidor
	e.Start(":8080")
}
