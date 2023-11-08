# book-collection <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

#### Esta é uma API para controle de livros

#### - Pesquise livros por (nome, autor, genero e ano de publicação)
#### - Favorite seus livros (Building)
#### - Deixe comentários (Building)

## Entidades
#### - Books (Livros do sistema)
#### - Users (Usuários do sistema)

## User Roles
#### - admin (usuários do sistema que gerencia a book collection)
#### - normal (Usuários normais do sistema)

## Installation

```bash
$ go mod tidy
```

## Running the app

Para conseguir rodar a aplicação sem erros, você precisará criar o arquivo .env com as seguintes variáveis

```
# Porta onde a API vai rodar
PORT=:3000

# URL de conexão com o banco de dados
DB_URL="host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"

# uma string que o JWT vai usar para gerar o token
SECRET="eyJhbGciOiJdfdfdssfreteryhfghjhgjhgj"
```

```bash
# first run migrate to create entites into DB
go run cmd/migrate/main.go

# to start the project
$ go run cmd/api/main.go
```
