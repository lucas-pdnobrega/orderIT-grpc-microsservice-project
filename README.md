# Go gRPC Microservices Example

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-v1.64-00D4B1.svg)](https://grpc.io/)
[![Docker](https://img.shields.io/badge/Docker-26.1-2496ED.svg)](https://www.docker.com/)

Guided microsservices project for PDIST. Contains two microservices:
- **Order Service**: Manages customer orders.
- **Payment Service**: Handles the payment processing for orders.
- **Shipping Service**: Handles the shipping processing for orders.

## Prerequisites

Ensure to have the following installed:
- Go (version 1.24.5 or newer)
- Docker and Docker Compose
- Git

## Running the Application

Here's the step by step (total of one :D) process for getting the application running.

### 1. Start the Database

The project uses a single docker container for all three services and the database images. Start it with Docker Compose:

```bash
docker-compose up --build
```

### 1. Testing Order's gRPC requests

#### This is a valid request to create an Order:

```bash
grpcurl -d '{"costumer_id": 1, "order_items": [{"product_code": "LAPTOP001", "quantity": 1, "unit_price": 350}]}' -plaintext localhost:3001 Order/Create
```

#### This is will result in a NOT FOUND exception due to not having "UNKNOWN" in inventory :

```bash
grpcurl -d '{"costumer_id": 1, "order_items": [{"product_code": "UNKNOWN", "quantity": 1, "unit_price": 350}]}' -plaintext localhost:3001 Order/Create
```

This is will result in a INVALID ARGUMENT exception due to having over 50 total items :

```bash
grpcurl -d '{"costumer_id": 1, "order_items": [{"product_code": "LAPTOP001", "quantity": 100, "unit_price": 350}]}' -plaintext localhost:3001 Order/Create
```

This is will result in a INVALID ARGUMENT exception due to having an over 1000 total payment amount :

```bash
grpcurl -d '{"costumer_id": 1, "order_items": [{"product_code": "LAPTOP001", "quantity": 10, "unit_price": 1350}]}' -plaintext localhost:3001 Order/Create
```

---

The application's running now. Cudos.
- The **Order Service** is available on port `3000`.
- The **Payment Service** is available on port `3001`.
- The **Shipping Service** is available on port `3002`.
- The **MySQL Database** is accessible on port `3310`.