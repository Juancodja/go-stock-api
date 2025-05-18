FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY .env ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-stock-pricer

# Etapa de ejecución
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/go-stock-pricer .

CMD ["source", ".env"]

CMD ["./go-stock-pricer"]
