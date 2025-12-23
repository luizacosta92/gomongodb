# API de Gestão de Gestantes

API RESTful desenvolvida em Go para cadastro e gerenciamento de usuárias gestantes, com persistência em MongoDB. Permite consultas por cidade, idade e data provável de parto (DPP).

---

##  Funcionalidades

-  Cadastro de gestantes
- Listagem de todas as gestantes
- Busca por cidade
- Busca por idade
- Busca por data provável de parto (DPP)

---

##  Tecnologias Utilizadas

| Tecnologia | Descrição |
|------------|-----------|
| **Go 1.21+** | Linguagem de programação |
| **Gin** | Framework web HTTP |
| **MongoDB** | Banco de dados NoSQL |
| **Docker** | Containerização |
| **Docker Compose** | Orquestração de containers |

---

## Estrutura do Projeto

```
exerciciomongodb/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── http/
│   │   └── routes/
│   │       └── router.go    # Configuração das rotas
│   ├── plataform/
│   │   └── mongo/
│   │       └── client.go    # Cliente de conexão MongoDB
│   └── user/
│       ├── handlers.go      # Handlers HTTP
│       ├── model.go         # Modelo de dados
│       ├── repository.go    # Interface do repositório
│       └── mongo_repository.go  # Implementação MongoDB
├── docker-compose.yml       # Configuração do MongoDB
├── .env                     # Variáveis de ambiente
├── go.mod
└── go.sum
```

---

## Como Executar

### Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Passo 1: Clone o repositório

```bash
git clone https://github.com/seu-usuario/exerciciomongodb.git
cd exerciciomongodb
```

### Passo 2: Inicie o MongoDB

```bash
docker-compose up -d
```

### Passo 3: Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
MONGO_URI=mongodb://workshop:pass123@localhost:27017/?authSource=admin
MONGO_DB=workshop
```

### Passo 4: Instale as dependências

```bash
go mod download
```

### Passo 5: Execute a aplicação

```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

---

## Endpoints da API

### Criar Usuária

```http
POST /users
Content-Type: application/json

{
  "name": "Marcela Avila",
  "whatsapp": "4899999-9999",
  "age": 32,
  "dpp": "2026-03-29",
  "city": "Florianópolis"
}
```

### Listar Todas

```http
GET /users
```

### Buscar por Cidade

```http
GET /users/city?city=Florianópolis
```

### Buscar por Idade

```http
GET /users/age?age=32
```

### Buscar por DPP

```http
GET /users/dpp?dpp=03-29-2026
```

---

##  Docker

O projeto inclui um `docker-compose.yml` para subir o MongoDB facilmente:

```bash
# Iniciar
docker-compose up -d

# Parar
docker-compose down

# Ver logs
docker logs workshop-mongo
```

---

## Variáveis de Ambiente

| Variável | Descrição | Exemplo |
|----------|-----------|---------|
| `MONGO_URI` | URI de conexão do MongoDB | `mongodb://workshop:pass123@localhost:27017/?authSource=admin` |
| `MONGO_DB` | Nome do banco de dados | `workshop` |
| `PORT` | Porta da API (opcional) | `8080` |

---

## Licença

Este projeto está sob a licença MIT.

---

Feito com 💜 e Go

