# Go Jobs API 🚀

Este é um projeto de estudo desenvolvido para consolidar aprendizados na
linguagem **Go (Golang)**. A aplicação consiste em uma API REST para
gerenciamento de oportunidades de emprego (vagas), focando em uma
arquitetura limpa e no uso de bibliotecas modernas do ecossistema Go.

---

## 🛠️ Tecnologias Utilizadas

- **Go (1.2x+)**
- **Gin Gonic**: Framework web de alta performance.
- **GORM**: ORM para persistência de dados.
- **SQLite**: Banco de dados relacional leve para desenvolvimento.
- **Swagger (Swaggo)**: Documentação automática da API.
- **Loggers Customizados**: Para melhor rastreabilidade de erros.

---

# Go Jobs API

API REST simples feita em Go para gerenciar oportunidades de emprego (vagas).

**Resumo:** aplicação com Gin + GORM usando SQLite para persistência local.

---

## Requisitos

- Go 1.18+ (ou compatível)
- `git`

---

## Como executar

1. Clone o repositório

```bash
git clone https://github.com/lucassoaresfr/go-api.git
cd go-api
```

2. Baixe dependências

```bash
go mod tidy
```

3. Execute a aplicação

```bash
go run main.go
```

A API escutará por padrão em `http://localhost:8080`.

O banco SQLite será criado em `./db/main.db` na primeira execução.

---

## Estrutura relevante

- `main.go` — inicializa logger, configurações e chama o roteador.
- `config/` — inicialização do DB (`sqlite.go`) e logger (`logger.go`).
- `router/` — define rotas e agrupa endpoints em `/api/v1`.
- `handler/` — handlers, validações e respostas padrão.
- `schemas/` — definição do modelo `Opening` usado pelo GORM.

---

## Endpoints (base: `/api/v1`)

- GET `/listOpenings`
  - Descrição: lista todas as vagas.
  - Query: nenhuma

- GET `/showOpening?id={ID}`
  - Descrição: retorna a vaga com o `id` informado.
  - Query: `id` (requerido)

- POST `/createOpening`
  - Descrição: cria uma nova vaga.
  - Body (JSON):
    - `role` (string, obrigatório)
    - `company` (string, obrigatório)
    - `location` (string, obrigatório)
    - `remote` (boolean, obrigatório)
    - `link` (string, obrigatório)
    - `salary` (number, obrigatório)

  Exemplo cURL:

  ```bash
  curl -X POST http://localhost:8080/api/v1/createOpening \
    -H "Content-Type: application/json" \
    -d '{"role":"Backend Engineer","company":"ACME","location":"Remote","remote":true,"link":"https://acme/jobs/1","salary":10000}'
  ```

- PUT `/updateOpening?id={ID}`
  - Descrição: atualiza campos da vaga com `id` informado.
  - Query: `id` (requerido)
  - Body (JSON): qualquer campo a ser atualizado (mesma estrutura de `create`), pelo menos um campo obrigatório pelo handler.

- DELETE `/deleteOpening?id={ID}`
  - Descrição: remove uma vaga pelo `id` (GORM soft delete disponível via `gorm.Model`).
  - Query: `id` (requerido)

---

## Observações de implementação

- Logger customizado em `config/logger.go` (métodos `Info`, `Error`, `Debug`).
- Banco de dados SQLite criado automaticamente em `./db/main.db` pela função `InitializeSQLite`.
- Validações de entrada estão em `handler/request.go`.
- Respostas padrão (erro/sucesso) em `handler/response.go`.

---

## Próximos passos sugeridos

- Adicionar testes unitários para handlers.
- Adicionar variáveis de ambiente para porta e caminho do DB.
- Documentar com Swagger (`swag`) se desejar interface interativa.

---

## Autor

Lucas Soares — projeto de estudos.
