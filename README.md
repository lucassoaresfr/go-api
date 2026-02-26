# Go Jobs API 🚀

Este é um projeto de estudo desenvolvido para consolidar aprendizados na
linguagem **Go (Golang)**. A aplicação consiste em uma API REST para
gerenciamento de oportunidades de emprego (vagas), focando em uma
arquitetura limpa e no uso de bibliotecas modernas do ecossistema Go.

------------------------------------------------------------------------

## 🛠️ Tecnologias Utilizadas

-   **Go (1.2x+)**
-   **Gin Gonic**: Framework web de alta performance.
-   **GORM**: ORM para persistência de dados.
-   **SQLite**: Banco de dados relacional leve para desenvolvimento.
-   **Swagger (Swaggo)**: Documentação automática da API.
-   **Loggers Customizados**: Para melhor rastreabilidade de erros.

------------------------------------------------------------------------

## 📂 Estrutura do Projeto

    /cmd        -> Ponto de entrada da aplicação (main.go)
    /handler    -> Lógica de controle das rotas (Create, List, Delete, etc)
    /schemas    -> Definição de modelos de dados e respostas (Structs)
    /db         -> Configuração de conexão e persistência SQLite
    /docs       -> Arquivos gerados automaticamente pelo Swagger

------------------------------------------------------------------------

## 🚀 Como Rodar o Projeto

### 1️⃣ Clone o repositório

``` bash
git clone https://github.com/lucassoaresfr/go-api.git
cd go-api
```

### 2️⃣ Instale as dependências

``` bash
go mod tidy
```

### 3️⃣ Gere a documentação do Swagger (opcional)

``` bash
swag init -g cmd/main.go
```

### 4️⃣ Execute a aplicação

``` bash
go run cmd/main.go
```

A API estará disponível em:

    http://localhost:8080

------------------------------------------------------------------------

## 📌 Endpoints Principais

  -----------------------------------------------------------------------
  Método             Endpoint                Descrição
  ------------------ ----------------------- ----------------------------
  GET                `/openings`             Lista todas as vagas (ignora
                                             as que sofreram soft delete)

  POST               `/opening`              Cria uma nova oportunidade

  DELETE             `/opening?id={ID}`      Remove uma vaga (suporta
                                             Soft Delete via GORM)

  GET                `/swagger/index.html`   Interface visual da
                                             documentação
  -----------------------------------------------------------------------

------------------------------------------------------------------------

## 📝 O que aprendi neste projeto

-   Manipulação de **Struct Tags** para mapeamento JSON.
-   Configuração de banco de dados relacional com **GORM**.
-   Implementação de **Soft Delete** e buscas customizadas.
-   Integração de documentação **Swagger** no Gin.
-   Gestão de variáveis de ambiente e caminhos de arquivos no Linux.

------------------------------------------------------------------------

## 👨‍💻 Autor

Desenvolvido por **Lucas Soares** como parte dos estudos em Back-end com
Go.