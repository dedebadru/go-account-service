FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o account-service ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/account-service .

COPY .env .env

CMD ["./account-service"]