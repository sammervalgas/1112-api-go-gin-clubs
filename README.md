# API Go Gin Clubs

Uma API RESTful desenvolvida em Go usando o framework Gin para gerenciamento de clubes e usuários.

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação
- [Gin](https://gin-gonic.com/) - Framework web
- [GORM](https://gorm.io/) - ORM para Go
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Docker](https://www.docker.com/) - Containerização

## 📋 Pré-requisitos

- Go 1.16+
- Docker e Docker Compose
- PostgreSQL (ou usar via Docker)

## 🔧 Instalação e Execução

1. Clone o repositório:
```bash
git clone https://github.com/sammervalgas/1112-api-go-gin-clubs.git

cd 1112-api-go-gin-clubs
```

2. Inicie o banco de dados com Docker:
```bash
cd infra/docker

docker-compose up -d
```

3. Execute a aplicação:
```bash
cd app

go run main.go
```

## 📚 Endpoints da API

### Clubs

- `GET /clubs` - Lista todos os clubes
- `GET /clubs/:pid` - Busca um clube pelo PID
- `POST /clubs` - Cria um novo clube
- `DELETE /clubs` - Remove um clube

### Users

- `GET /users` - Lista todos os usuários
- `GET /users/:id` - Busca um usuário pelo ID
- `POST /users` - Cria um novo usuário
- `DELETE /users/:id` - Remove um usuário
- `PATCH /users/:id` - Atualiza um usuário

## 📝 Exemplos de Uso

### Criar um novo usuário
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "cpf": "12345678901",
    "birthday": "1990-01-01T00:00:00Z"
  }'
```

### Criar um novo clube
```bash
curl -X POST http://localhost:8080/clubs \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Club Example",
    "address": "123 Main St",
    "pid": "3d01b30c-ad24-4f94-a945-26f4e2b95054",
    "website": "example.com"
  }'
```

## 🗄️ Estrutura do Banco de Dados

A estrutura de banco de dados esta sendo criada automaticamente pelo GORM.

### Tabela Clubs
```sql
CREATE TABLE clubs (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    pid UUID NOT NULL,
    website VARCHAR(255) NOT NULL
);
```

### Tabela Users
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    birthday TIMESTAMP NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

## 🔧 Acessando o PGAdmin

O PGAdmin é uma ferramenta de administração e desenvolvimento para PostgreSQL. Para acessar:

1. Abra seu navegador e acesse: http://localhost:54321

2. Use as seguintes credenciais:
   - Email: go@devbean.com.br 
   - Senha: 12356

3. Para conectar ao banco de dados PostgreSQL:
   - Host: postgres
   - Porta: 5432
   - Usuário: root
   - Senha: root
   - Banco de dados: root



## ✒️ Autores

* **Sammer Valgas**
