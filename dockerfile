# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app

# copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# copia o restante do código
COPY . .

# compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o cesjb ./cmd/app/main.go

# Run stage — imagem menor para produção
FROM alpine:latest

WORKDIR /app

# copia o binário compilado
COPY --from=builder /app/cesjb .

# copia o .env
COPY cmd/app/.env .

EXPOSE 8088

CMD ["./cesjb"]