## PARA EXECUTAR A API SEM O DOCKER BASTA SEGUIR OS SEGUINTES PASSOS

- Baixe o postgres
- Crie um banco chamado estoque
- Altere no arquivo config.go os enviroments
- crie uma table com a seguinte query
CREATE TABLE IF NOT EXISTS public.produtos
(
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    preco NUMERIC(10,2) NOT NULL,
    quantidade INTEGER NOT NULL,
    categoria VARCHAR(50),
    desconto NUMERIC(5,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
- Clone o repositorio com git clone <nome do repositorio>
- Entre no arquivo com cd estoque-api
- Instale a linguagem go no seu dispositivo
- Execute a API com o comando go run main.go 

## PARA EXECUTAR A API COM O DOCKER BASTA SEGUIR OS SEGUINTES PASSOS:

- git clone <nome do repositorio>
- cd estoque-api
- docker compose up --build

## PARA REALIZAR REQUISIÇÕES DIRETO NO WINDOWS POWERSHELL

Gera o token - $tokenResponse = Invoke-RestMethod -Uri "http://localhost:8080/login" -Method Post

Armazena o token em uma variavel - $token = $tokenResponse.token

Exibe o token gerado - Write-Host "Token JWT gerado: $token"

GET - Listar produtos - Invoke-RestMethod -Uri "http://localhost:8080/produtos" -Method Get -Headers @{Authorization=("Bearer " + $token)}

POST - criar novo prduto - Invoke-RestMethod -Uri "http://localhost:8080/produtos" -Method Post -Headers @{Authorization=("Bearer " + $token)} -Body (@{
    nome = "Produto Exemplo"
    descricao = "Descrição do produto exemplo"
    preco = 150.00
    quantidade = 10
    categoria = "Eletrônicos"
    desconto = 5.0
} | ConvertTo-Json) -ContentType "application/json"

PUT - Atualizar produto - Invoke-RestMethod -Uri "http://localhost:8080/produtos/:id" -Method Put -Headers @{Authorization=("Bearer " + $token)} -Body (@{
    nome = "Produto Atualizado"
    descricao = "Nova descrição"
    preco = 120.00
    quantidade = 5
    categoria = "Eletrônicos"
    desconto = 10.0
} | ConvertTo-Json) -ContentType "application/json"

DELETE - excluir produto - Invoke-RestMethod -Uri "http://localhost:8080/produtos/:id" -Method Delete -Headers @{Authorization=("Bearer " + $token)}
