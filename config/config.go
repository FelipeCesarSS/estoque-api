package config

import (
	"fmt"
	"log"

	"github.com/FelipeCesarSS/estoque-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=db user=postgres password=admin dbname=estoque port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados!")
	}

	DB = database
	fmt.Println("Conexão com o banco de dados realizada com sucesso!")
}

func Migrate() {
	err := DB.AutoMigrate(&models.Produto{})
	if err != nil {
		log.Fatal("Falha ao migrar o banco de dados:", err)
	}
}
