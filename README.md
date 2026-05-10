# Gobid

API em Go para um sistema de leiloes online com cadastro de usuarios, criacao de produtos e lances em tempo real via WebSocket.

## O que este projeto faz

- Cadastro de usuarios com `signup`, `login` e `logout`
- Criacao de produtos com data final do leilao
- Participacao em leiloes em tempo real
- Envio de lances pelo WebSocket
- Valida validacoes de entrada antes de salvar no banco
- Usa sessoes para manter usuario autenticado

## Tecnologias usadas

- Go
- Chi para rotas e middlewares
- PostgreSQL
- pgx para acesso ao banco
- sqlc para gerar queries SQL tipadas
- scs para gerenciamento de sessoes
- gorilla/websocket para comunicacao em tempo real
- gorilla/csrf preparado no projeto
- bcrypt para hash de senha

## Estrutura principal

- `cmd/api`: inicializacao da API
- `internal/api`: handlers e rotas HTTP
- `internal/services`: regras de negocio
- `internal/usecase`: validacao dos payloads
- `internal/store/pgstore`: queries e modelos gerados pelo `sqlc`

## Rotas principais

- `POST /api/v1/users/signup`
- `POST /api/v1/users/login`
- `POST /api/v1/users/logout`
- `POST /api/v1/products/`
- `GET /api/v1/products/ws/subscribe/{product_id}`

## Banco de dados

O projeto usa PostgreSQL com estas tabelas principais:

- `users`
- `sessions`
- `products`
- `bids`

## Como executar

1. Configure o arquivo `.env` com as variaveis do banco.
2. Suba o PostgreSQL com o `docker-compose.yml`.
3. Rode a aplicacao em `cmd/api`.

## Variaveis de ambiente

- `GOBID_DATABASE_PORT`
- `GOBID_DATABASE_NAME`
- `GOBID_DATABASE_USER`
- `GOBID_DATABASE_PASSWORD`
- `GOBID_DATABASE_HOST`

## Observacao

O projeto implementa um leilao com sala em memoria para cada produto, e o fechamento do leilao acontece automaticamente quando o prazo definido para o produto termina.
