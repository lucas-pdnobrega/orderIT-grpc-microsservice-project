package ports

import "shipping/internal/application/core/domain"

type APIPort interface {
	CreateShipping(domain.Shipping) (domain.Shipping, error)
}
