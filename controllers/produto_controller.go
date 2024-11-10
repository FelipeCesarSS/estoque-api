package controllers

import (
	"net/http"
	"strconv"

	"github.com/FelipeCesarSS/estoque-api/models"
	"github.com/FelipeCesarSS/estoque-api/services"
	"github.com/labstack/echo/v4"
)

// ListarProdutos handler para listar todos os produtos
func ListarProdutos(c echo.Context) error {
	produtos, err := services.ListarTodosProdutos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, produtos)
}

// CriarProduto handler para adicionar um novo produto
func CriarProduto(c echo.Context) error {
	produto := new(models.Produto)
	if err := c.Bind(produto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dados inválidos"})
	}

	if err := services.CriarProduto(produto); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, produto)
}

// AtualizarProduto handler para atualizar um produto existente
func AtualizarProduto(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	dadosAtualizados := new(models.Produto)
	if err := c.Bind(dadosAtualizados); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dados inválidos"})
	}

	produto, err := services.AtualizarProduto(uint(id), dadosAtualizados)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, produto)
}

// DeletarProduto handler para excluir um produto logicamente
func DeletarProduto(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	if err := services.DeletarProduto(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
