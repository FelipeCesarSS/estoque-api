package services

import (
	"errors"

	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/models"
)

// ListarTodosProdutos retorna todos os produtos disponíveis
func ListarTodosProdutos() ([]models.Produto, error) {
	var produtos []models.Produto
	// Consulta no banco com soft delete (onde deleted_at é NULL)
	result := config.DB.Where("deleted_at IS NULL").Find(&produtos)
	if result.Error != nil {
		return nil, result.Error // Retorna o erro se falhar
	}
	return produtos, nil // Retorna a lista de produtos
}

// ObterProdutoPorID retorna um produto específico pelo ID
func ObterProdutoPorID(id uint) (*models.Produto, error) {
	var produto models.Produto
	// Faz a busca no banco considerando o soft delete
	result := config.DB.First(&produto, "id = ? AND deleted_at IS NULL", id)
	if result.Error != nil {
		return nil, errors.New("Produto não encontrado") // Caso não encontre, retorna erro
	}
	return &produto, nil // Retorna o produto encontrado
}

// CriarProduto cria um novo produto no banco de dados
func CriarProduto(produto *models.Produto) error {
	result := config.DB.Create(produto)
	if result.Error != nil {
		return result.Error // Retorna erro se não conseguir criar
	}
	return nil // Retorna nil se criar com sucesso
}

// AtualizarProduto atualiza um produto existente
func AtualizarProduto(id uint, dadosAtualizados *models.Produto) (*models.Produto, error) {
	produtoExistente, err := ObterProdutoPorID(id)
	if err != nil {
		return nil, err // Caso o produto não exista, retorna erro
	}

	// Atualiza os campos do produto com os dados recebidos
	produtoExistente.Nome = dadosAtualizados.Nome
	produtoExistente.Descricao = dadosAtualizados.Descricao
	produtoExistente.Preco = dadosAtualizados.Preco
	produtoExistente.Quantidade = dadosAtualizados.Quantidade
	produtoExistente.Categoria = dadosAtualizados.Categoria
	produtoExistente.Desconto = dadosAtualizados.Desconto

	// Salva as mudanças no banco
	result := config.DB.Save(produtoExistente)
	if result.Error != nil {
		return nil, result.Error // Retorna erro se falhar
	}
	return produtoExistente, nil // Retorna o produto atualizado
}

// DeletarProduto realiza uma exclusão lógica de um produto (soft delete)
func DeletarProduto(id uint) error {
	produtoExistente, err := ObterProdutoPorID(id)
	if err != nil {
		return err // Caso não encontre o produto, retorna erro
	}

	// Realiza o soft delete atualizando o campo deleted_at com a data atual
	result := config.DB.Model(produtoExistente).Update("deleted_at", "NOW()")
	if result.Error != nil {
		return result.Error // Retorna erro se não conseguir excluir
	}
	return nil // Retorna nil se deletar com sucesso
}
