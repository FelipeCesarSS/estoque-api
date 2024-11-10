# Manipulaçao da API

Gerar token - Invoke-RestMethod -Uri "http://localhost:8080/login" -Method Post

Armazenar o token em uma variável - $token = "<coloque_aqui_o_token_gerado>"

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
