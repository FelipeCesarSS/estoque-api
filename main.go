package main

import (
	"net/http"
	"time"

	"github.com/FelipeCesarSS/estoque-api/config"
	"github.com/FelipeCesarSS/estoque-api/routes"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Chave secreta para assinar o JWT
var secretKey = []byte("seu_segredo")

// Função para gerar o token JWT
func gerarToken() (string, error) {
	claims := jwt.MapClaims{
		"sub": "usuario1",
		"exp": time.Now().Add(time.Hour * 1).Unix(), //Expira em 7 dias
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Função para login - gera o token JWT
func login(c echo.Context) error {
	token, err := gerarToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erro ao gerar o token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func main() {
	config.ConnectDatabase()

	e := echo.New()

	// Rota para gerar o token
	e.POST("/login", login)

	// Configurar as rotas dos produtos
	routes.InitRoutes(e)

	// Iniciar servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}
