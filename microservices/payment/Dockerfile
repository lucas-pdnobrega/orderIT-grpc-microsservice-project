FROM golang:1.24.5-alpine AS builder

WORKDIR /app/microservices/payment
COPY ./microservices/payment /app/microservices/payment
COPY ./microservices-proto/ /app/microservices-proto/

RUN go mod tidy
RUN go build -o payment-service ./cmd/main.go

EXPOSE 3003

CMD ["/app/microservices/payment/payment-service"]