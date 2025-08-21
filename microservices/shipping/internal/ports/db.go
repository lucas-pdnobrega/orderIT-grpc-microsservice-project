package ports

import "shipping/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Shipping, error)
	Save(*domain.Shipping) error
}
