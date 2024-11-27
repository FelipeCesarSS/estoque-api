package services

import (
	"errors"
	"time"

	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/models"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("seu_segredo")

func GerarToken() (string, error) {
	claims := jwt.MapClaims{
		"sub": "usuario1",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ListarTodosProdutos() ([]models.Produto, error) {
	var produtos []models.Produto
	result := config.DB.Where("deleted_at IS NULL").Find(&produtos)
	if result.Error != nil {
		return nil, result.Error
	}
	return produtos, nil
}

func ObterProdutoPorID(id uint) (*models.Produto, error) {
	var produto models.Produto
	result := config.DB.First(&produto, "id = ? AND deleted_at IS NULL", id)
	if result.Error != nil {
		return nil, errors.New("Produto não encontrado")
	}
	return &produto, nil
}

func CriarProduto(produto *models.Produto) error {
	result := config.DB.Create(produto)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AtualizarProduto(id uint, dadosAtualizados *models.Produto) (*models.Produto, error) {
	produtoExistente, err := ObterProdutoPorID(id)
	if err != nil {
		return nil, err
	}

	produtoExistente.Nome = dadosAtualizados.Nome
	produtoExistente.Descricao = dadosAtualizados.Descricao
	produtoExistente.Preco = dadosAtualizados.Preco
	produtoExistente.Quantidade = dadosAtualizados.Quantidade
	produtoExistente.Categoria = dadosAtualizados.Categoria
	produtoExistente.Desconto = dadosAtualizados.Desconto

	result := config.DB.Save(produtoExistente)
	if result.Error != nil {
		return nil, result.Error
	}
	return produtoExistente, nil
}

func DeletarProduto(id uint) error {
	produtoExistente, err := ObterProdutoPorID(id)
	if err != nil {
		return err
	}

	result := config.DB.Model(produtoExistente).Update("deleted_at", "NOW()")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
