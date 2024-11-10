package models

import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
	Categoria  string  `json:"categoria"`
	Desconto   float64 `json:"desconto"`
}
