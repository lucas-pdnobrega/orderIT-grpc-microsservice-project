package main

import (
	"log"

	"shipping/config"
	"shipping/internal/adapters/db"
	"shipping/internal/adapters/grpc"
	"shipping/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error : %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	log.Println("Shipping service's gonna run...")
	grpcAdapter.Run()
}
