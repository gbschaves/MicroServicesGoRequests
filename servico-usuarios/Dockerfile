# Etapa de Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copia dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila a aplicação (garante um binário estático)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa Final (Imagem leve)
FROM alpine:latest

WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/main .
# Copia a pasta docs para o Swagger funcionar (se necessário)
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]