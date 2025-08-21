package api

import (
	"shipping/internal/application/core/domain"
	"shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreateShipping(shipping domain.Shipping) (domain.Shipping, error) {
	if err := a.db.Save(&shipping); err != nil {
		return domain.Shipping{}, err
	}

	return shipping, nil
}
