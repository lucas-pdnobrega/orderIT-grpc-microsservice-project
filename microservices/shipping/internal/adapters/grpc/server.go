package grpc

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"shipping/config"
	"shipping/internal/application/core/domain"
	"shipping/internal/ports"

	"github.com/ruandg/microservices-proto/golang/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api  ports.APIPort
	port int
	shipping.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating order...")

	var shippingItems []domain.ShippingItem
	for _, shippingItem := range request.ShippingItems {
		shippingItems = append(shippingItems, domain.ShippingItem{
			ProductCode: shippingItem.ProductCode,
			Quantity:    shippingItem.Quantity,
		})
	}

	newShipping := domain.NewShipping(int64(request.OrderId), shippingItems)
	result, err := a.api.CreateShipping(newShipping)
	code := status.Code(err)
	if code == codes.InvalidArgument {
		return nil, err
	} else if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to ship! %v ", err)).Err()
	}
	return &shipping.CreateShippingResponse{DeliverySpan: int32(result.DeliverySpan)}, nil
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}
