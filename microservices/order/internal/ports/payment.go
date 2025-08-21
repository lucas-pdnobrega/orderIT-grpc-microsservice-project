package ports

import "order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order domain.Order) error
}
