# CESJB — Sistema de Gerenciamento de Associados

Sistema desenvolvido para o **Centro Espírita São João Batista**, com foco no gerenciamento de associados e controle de mensalidades.

---

## Tecnologias Utilizadas

### Backend
- **Go** — linguagem principal
- **net/http** — biblioteca padrão para construção da API REST
- **PostgreSQL** — banco de dados relacional
- **pgx/v5** — driver PostgreSQL com pool de conexões (`pgxpool`)
- **golang-jwt/jwt/v5** — autenticação via JWT
- **godotenv** — gerenciamento de variáveis de ambiente
- **bcrypt** — hash de senhas

### Frontend
- **HTML5, CSS3 e JavaScript puro** — sem frameworks
- Interface responsiva com sidebar, modais e paginação
- Integração com a API via `fetch`

### Infraestrutura
- **Docker & Docker Compose** — banco de dados em container local

---

## Arquitetura

O projeto segue uma arquitetura em camadas inspirada no padrão **Clean Architecture**:

```
cmd/app/         → ponto de entrada da aplicação
api/             → inicialização das rotas e interfaces
handlers/        → recebimento e validação das requisições HTTP
domain/
  entities/      → entidades do domínio
  service/       → regras de negócio
  errors.go      → erros de domínio centralizados
repository/      → acesso ao banco de dados (PostgreSQL)
middlewares/     → autenticação JWT, logger e CORS
authentication/  → geração e validação de tokens
config/          → carregamento de variáveis de ambiente
types_/          → tipo customizado DateOnly
```

---

## Funcionalidades

### Associados
- Cadastro com validação de CPF (dígitos verificadores), e-mail e campos obrigatórios
- Edição de dados
- Desativação e reativação
- Listagem de associados ativos e inativos
- Busca por nome, CPF ou ID

### Pagamentos
- Registro de mensalidades
- Listagem de pagamentos por mês
- Total arrecadado por competência
- Listagem de inadimplentes por mês
- Histórico de pagamentos por associado

### Admin
- Cadastro de administrador
- Login com geração de token JWT (validade de 6 horas)

---

## Rotas da API

### Associados
| Método | Rota | Descrição | Auth |
|--------|------|-----------|------|
| `POST` | `/associate` | Cadastrar associado | ✅ |
| `GET` | `/associates` | Listar associados ativos | ✅ |
| `GET` | `/associates/inactive` | Listar associados inativos | ✅ |
| `GET` | `/associate/id/{id}` | Buscar por ID | ✅ |
| `GET` | `/associate/cpf/{cpf}` | Buscar por CPF | ✅ |
| `GET` | `/associate/name/{name}` | Buscar por nome | ✅ |
| `PUT` | `/associate/{id}` | Atualizar associado | ✅ |

### Pagamentos
| Método | Rota | Descrição | Auth |
|--------|------|-----------|------|
| `POST` | `/payment` | Registrar pagamento | ✅ |
| `GET` | `/payments?competence=yyyy-mm-dd` | Listar pagamentos do mês | ✅ |
| `GET` | `/payments/month?competence=yyyy-mm-dd` | Total arrecadado no mês | ✅ |
| `GET` | `/payments/defaulters?competence=yyyy-mm-dd` | Inadimplentes do mês | ✅ |
| `GET` | `/payments/associate/{id}` | Pagamentos por associado | ✅ |

### Admin
| Método | Rota | Descrição | Auth |
|--------|------|-----------|------|
| `POST` | `/admin` | Cadastrar administrador | ❌ |
| `POST` | `/login` | Login e geração de token | ❌ |

---

## Como Executar

### Pré-requisitos
- Go 1.22+
- PostgreSQL
- Docker (opcional)

### Variáveis de Ambiente

Crie um arquivo `.env` em `cmd/app/`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=cesjb_db
API_PORT=8088
SECRET_KEY=sua_chave_secreta
```

### Rodando a API

```bash
cd cmd/app
go run main.go
```

---

## Frontend

O frontend é uma aplicação web estática localizada na pasta `frontend/`, composta por:

```
frontend/
├── index.html              → tela de login
├── css/                    → estilos por página
├── js/                     → lógica por página
└── pages/
    ├── dashboard.html      → tela inicial com resumo
    ├── associates.html     → gestão de associados
    ├── associate_detail.html → ficha do associado
    └── payments.html       → gestão de pagamentos
```

Para utilizar, abra o `index.html` diretamente no navegador com a API rodando.
