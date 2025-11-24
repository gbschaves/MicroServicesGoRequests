# Etapa de Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Instala git (caso seja necessário para baixar dependências)
RUN apk add --no-cache git

# Copia dependências
COPY go.mod go.sum ./
# Tenta baixar o que já está no go.sum antigo
RUN go mod download

# Copia o código fonte
COPY . .

# --- CORREÇÃO: Gera o go.sum correto incluindo o lib/pq ---
RUN go mod tidy
# ----------------------------------------------------------

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa Final (Imagem leve)
FROM alpine:latest

WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/main .
# Copia a pasta docs para o Swagger funcionar
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]