package routes

import (
	"github.com/FelipeCesarSS/estoque-api/controllers"
	"github.com/FelipeCesarSS/estoque-api/middleware"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.POST("/login", controllers.Login)

	e.GET("/produtos", controllers.ListarProdutos)

	produtos := e.Group("/produtos")
	produtos.Use(middleware.JWTMiddleware)

	produtos.POST("", controllers.CriarProduto)

	produtos.PUT("/:id", controllers.AtualizarProduto)

	produtos.DELETE("/:id", controllers.DeletarProduto)
}
