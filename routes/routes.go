package routes

import (
	"github.com/FelipeCesarSS/estoque-api/controllers"
	"github.com/FelipeCesarSS/estoque-api/middleware"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/produtos", controllers.ListarProdutos, middleware.JWTMiddleware())
	e.POST("/produtos", controllers.CriarProduto, middleware.JWTMiddleware())
	e.PUT("/produtos/:id", controllers.AtualizarProduto, middleware.JWTMiddleware())
	e.DELETE("/produtos/:id", controllers.DeletarProduto, middleware.JWTMiddleware())
}
