# book-collection <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

#### Esta é uma API para controle de livros

#### - Pesquisar livros por (nome, autor, genero e ano de publicação)

## Entidades
### - Books (Livros do sistema)
### - Users (Usuários do sistema)
#### -- admin (Usuário que gerencia os livros)
#### -- normal (Usuário que aluga livros)
 
## Installation

```bash
$ go mod tidy
```

## Running the app

Para conseguir rodar a aplicação sem erros, você precisará criar o arquivo .env com as seguintes variáveis

```
# Porta onde a API vai rodar
PORT=3000

# URL de conexão com o banco de dados
DB_URL="host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
```

```bash
# start
$ go run main.go
```
