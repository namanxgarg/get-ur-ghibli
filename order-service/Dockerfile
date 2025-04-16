FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service ./cmd/main.go

EXPOSE 8084
CMD ["./order-service"]
