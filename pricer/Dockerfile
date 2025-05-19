FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-stock-pricer

# Etapa de ejecución
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/go-stock-pricer .

RUN chmod +x ./go-stock-pricer


CMD ["./go-stock-pricer"]
