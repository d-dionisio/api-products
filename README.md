# API Products

API REST simples para gerenciamento de produtos, desenvolvida em Go com o pacote padrão `net/http`.

O projeto implementa operações de criação, consulta, atualização e exclusão (CRUD). Os produtos são armazenados em memória, portanto os dados voltam ao estado inicial sempre que a aplicação é reiniciada.

## Tecnologias

- Go 1.26.3
- `net/http`
- `encoding/json`

## Como executar

### Pré-requisitos

Tenha o [Go](https://go.dev/dl/) instalado em sua máquina.

### Iniciando a API

```bash
go run .
```

O servidor estará disponível em:

```text
http://localhost:8000
```

## Diagrama da API

```mermaid
flowchart LR
    Client[Cliente HTTP] --> Router{Rota e método}

    Router -->|GET /products| List[Listar produtos]
    Router -->|POST /products| Create[Criar produto]
    Router -->|GET /products/id| Get[Buscar produto]
    Router -->|PUT /products/id| Update[Atualizar produto]
    Router -->|DELETE /products/id| Delete[Excluir produto]

    List --> Store[(Produtos em memória)]
    Create --> Store
    Get --> Store
    Update --> Store
    Delete --> Store

    Store --> Response[Resposta HTTP em JSON]
    Response --> Client
```

## Modelo de produto

```json
{
  "id": 1,
  "name": "Monitor",
  "price": 700
}
```

| Campo | Tipo | Descrição |
| --- | --- | --- |
| `id` | inteiro | Identificador gerado pela API |
| `name` | string | Nome do produto |
| `price` | decimal | Preço do produto |

## Endpoints

| Método | Rota | Descrição | Resposta de sucesso |
| --- | --- | --- | --- |
| `GET` | `/products` | Lista todos os produtos | `200 OK` |
| `POST` | `/products` | Cria um produto | `201 Created` |
| `GET` | `/products/{id}` | Busca um produto pelo ID | `200 OK` |
| `PUT` | `/products/{id}` | Substitui os dados de um produto | `200 OK` |
| `DELETE` | `/products/{id}` | Exclui um produto | `204 No Content` |

IDs inválidos retornam `400 Bad Request`. Quando o produto não existe, a API retorna `404 Not Found`.

## Exemplos de uso

### Listar produtos

```bash
curl http://localhost:8000/products
```

### Criar um produto

```bash
curl -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Mouse","price":80}'
```

### Buscar um produto

```bash
curl http://localhost:8000/products/1
```

### Atualizar um produto

```bash
curl -X PUT http://localhost:8000/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Monitor Full HD","price":1200}'
```

### Excluir um produto

```bash
curl -X DELETE http://localhost:8000/products/1
```

As mesmas requisições também estão disponíveis no arquivo [`request.http`](request.http), que pode ser executado por extensões como REST Client no VS Code.

## Estrutura do projeto

```text
api-products/
├── go.mod
├── main.go
├── README.md
└── request.http
```

## Observações

- Não há banco de dados: todas as alterações são perdidas ao encerrar o servidor.
- A API começa com dois produtos de exemplo: Monitor e Teclado.
- O campo `id` enviado pelo cliente é ignorado nas criações e definido automaticamente pela API.
