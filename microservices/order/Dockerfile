FROM golang:1.24.5-alpine AS builder

WORKDIR /app/microservices/order
COPY ./microservices/order /app/microservices/order
COPY ./microservices-proto/ /app/microservices-proto/

RUN go mod tidy
RUN go build -o order-service ./cmd/main.go

EXPOSE 3001

CMD ["/app/microservices/order/order-service"]