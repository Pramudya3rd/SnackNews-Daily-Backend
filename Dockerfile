FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o news-shared-service ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/news-shared-service .

CMD ["./news-shared-service"]