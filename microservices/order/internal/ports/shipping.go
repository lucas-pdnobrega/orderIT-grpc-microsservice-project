package ports

import "order/internal/application/core/domain"

type ShippingPort interface {
	CreateShipping(order domain.Order) (error)
}
