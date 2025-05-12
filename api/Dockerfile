FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-stock-api

# Etapa de ejecución
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/go-stock-api .

EXPOSE 8080

CMD ["./go-stock-api"]
