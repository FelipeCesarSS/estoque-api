package routes

import (
	"github.com/FelipeCesarSS/estoque-api/controllers"
	"github.com/FelipeCesarSS/estoque-api/middleware"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Rota pública: listar produtos
	e.GET("/produtos", controllers.ListarProdutos)

	// Grupo de rotas protegidas por JWT
	produtos := e.Group("/produtos")
	produtos.Use(middleware.JWTMiddleware)

	// Rota para criar produto
	produtos.POST("", controllers.CriarProduto)

	// Rota para atualizar produto
	produtos.PUT("/:id", controllers.AtualizarProduto)

	// Rota para deletar produto
	produtos.DELETE("/:id", controllers.DeletarProduto)
}
