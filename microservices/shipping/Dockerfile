FROM golang:1.24.5-alpine AS builder

WORKDIR /app/microservices/shipping
COPY ./microservices/shipping /app/microservices/shipping
COPY ./microservices-proto/ /app/microservices-proto/

RUN go mod tidy
RUN go build -o shipping-service ./cmd/main.go

EXPOSE 3003

CMD ["/app/microservices/shipping/shipping-service"]