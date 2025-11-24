# Etapa de Build - MUDAMOS DE 1.21 PARA 1.22
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa Final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
# A linha abaixo é opcional, só se tiver a pasta docs
# COPY --from=builder /app/docs ./docs

EXPOSE 8080 8081 8082 8083

CMD ["./main"]