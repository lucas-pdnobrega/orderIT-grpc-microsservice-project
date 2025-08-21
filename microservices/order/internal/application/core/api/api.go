package api

import (
	"order/internal/application/core/domain"
	"order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	var totalQuantity int32
	for _, orderItem := range order.OrderItems {
		totalQuantity += orderItem.Quantity

		_, err := a.db.FindInventoryItemByProductCode(orderItem.ProductCode)
		if err != nil {
			return domain.Order{}, status.Errorf(
				codes.NotFound,
				"Product with code %s not found",
				orderItem.ProductCode,
			)
		}
	}

	if totalQuantity > 50 {
		return domain.Order{}, status.Errorf(
			codes.InvalidArgument,
			"Order's total item quantity is %d, it cannot exceed 50",
			totalQuantity,
		)
	}
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	shippingErr := a.shipping.CreateShipping(order)
	if shippingErr != nil {
		return domain.Order{}, shippingErr
	}

	return order, nil
}
