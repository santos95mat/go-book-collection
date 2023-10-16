# book-collection <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

#### Esta é uma API para armazenar dados de livros

#### - Pesquisar livros por (nome, autor, tema)

## Entidades
#### - Books (Livros do sistema)
#### - Autor (Autores do sistema, relação one-to-many com os Livros)
 
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
