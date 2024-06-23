# Loja com Golang

Este é um projeto de uma aplicação Web em Go para uma loja virtual simples, utilizando MySQL como banco de dados e HTML com Bootstrap para as interfaces.

## Requisitos

- Go 1.16 ou superior
- MySQL

## Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/pedrolessa-dev/loja-golang.git
    ```

2. Instale as dependências:

    ```bash
    go mod tidy
    ```

3. Configure as variáveis de ambiente para a conexão com o banco de dados MySQL:

    ```bash
    export MYSQL_USER=seu_usuario_mysql
    export MYSQL_PASSWORD=sua_senha_mysql
    ```

4. Execute a aplicação:

    ```bash
    go run main.go
    ```

5. Acesse a aplicação em ***<http://localhost:8000>***.

## Funcionalidades

- Listagem de produtos
- Adição de novos produtos
- Edição de produtos existentes
- Exclusão de produtos

## Estrutura do Projeto

- `main.go`: Arquivo principal que inicia a aplicação e carrega as rotas.
- `routes/routes.go`: Define as rotas da aplicação.
- `controllers/produtos.go`: Contém as funções que controlam o fluxo da aplicação.
- `models/produtos.go`: Define a estrutura dos dados e contém as funções relacionadas ao banco de dados.
- `db/db.go`: Arquivo responsável por estabelecer a conexão com o banco de dados.
