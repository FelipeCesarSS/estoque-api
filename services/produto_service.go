package services

import (
	"errors"

	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/models"
)

// ListarTodosProdutos retorna todos os produtos disponíveis
func ListarTodosProdutos() ([]models.Produto, error) {
	var produtos []models.Produto
	result := config.DB.Where("deleted_at IS NULL").Find(&produtos)
	if result.Error != nil {
		return nil, result.Error
	}
	return produtos, nil
}

// ObterProdutoPorID retorna um produto específico pelo ID
func ObterProdutoPorID(id uint) (*models.Produto, error) {
	var produto models.Produto
	result := config.DB.First(&produto, "id = ? AND deleted_at IS NULL", id)
	if result.Error != nil {
		return nil, errors.New("Produto não encontrado")
	}
	return &produto, nil
}

// CriarProduto cria um novo produto no banco de dados
func CriarProduto(produto *models.Produto) error {
	result := config.DB.Create(produto)
	return result.Error
}

// AtualizarProduto atualiza um produto existente
func AtualizarProduto(id uint, dadosAtualizados *models.Produto) (*models.Produto, error) {
	produtoExistente, err := ObterProdutoPorID(id)
	if err != nil {
		return nil, err
	}

	// Atualiza os campos necessários
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

// DeletarProduto realiza uma exclusão lógica de um produto (soft delete)
func DeletarProduto(id uint) error {
	produtoExistente, err := ObterProdutoPorID(id)
	if err != nil {
		return err
	}

	result := config.DB.Model(produtoExistente).Update("deleted_at", "NOW()")
	return result.Error
}
