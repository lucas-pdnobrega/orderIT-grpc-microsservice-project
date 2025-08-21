package main

import (
	"log"

	"order/config"
	"order/internal/adapters/db"
	"order/internal/adapters/grpc"
	"order/internal/adapters/payment"
	"order/internal/adapters/shipping"
	"order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error : %v", err)
	}
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error : %v", err)
	}
	shippingAdapter, err := shipping_adapter.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error : %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	log.Println("Order service's gonna run...")
	grpcAdapter.Run()
}
