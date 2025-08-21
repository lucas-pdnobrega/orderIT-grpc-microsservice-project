package shipping_adapter

import (
	"context"
	"log"
	"time"

	"order/internal/application/core/domain"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/ruandg/microservices-proto/golang/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	shipping shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(2*time.Second)),
		)),
	)
	opts = append(
		opts, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	conn, err := grpc.Dial(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := shipping.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) CreateShipping(order domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var shippingItems []*shipping.ShippingItem
	for _, item := range order.OrderItems {
		shippingItems = append(shippingItems, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    int32(item.Quantity),
		})
	}

	_, err := a.shipping.Create(ctx, &shipping.
		CreateShippingRequest{
		OrderId:    order.ID,
		ShippingItems: shippingItems,
	})

	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.DeadlineExceeded {
			log.Printf("Timeout: deadline exceeded for Payment.Charge for order %d", order.ID)
		}
		return err
	}

	return nil
}
