package db

import (
	"fmt"

	"shipping/internal/application/core/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	ShippingID int64
	ShippingItems []ShippingItem
	DeliverySpan int32
}

type ShippingItem struct {
	gorm.Model
	ProductCode string
	Quantity    int32
	ShippingID uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Shipping{}, ShippingItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id string) (domain.Shipping, error) {
	var shippingEntity Shipping
	res := a.db.First(&shippingEntity, id)
	var shippingItems []domain.ShippingItem
	for _, shippingItem := range shippingEntity.ShippingItems {
		shippingItems = append(shippingItems, domain.ShippingItem{
			ProductCode: shippingItem.ProductCode,
			Quantity:    shippingItem.Quantity,
		})
	}
	shipping := domain.Shipping{
		ID:         int64(shippingEntity.ShippingID),
		ShippingItems: shippingItems,
		CreatedAt:  shippingEntity.CreatedAt.UnixNano(),
	}
	return shipping, res.Error
}

func (a Adapter) Save(shipping *domain.Shipping) error {
	var shippingItems []ShippingItem
	for _, shippingItem := range shipping.ShippingItems {
		shippingItems = append(shippingItems, ShippingItem{
			ProductCode: shippingItem.ProductCode,
			Quantity:    shippingItem.Quantity,
		})
	}
	shippingModel := Shipping{
		ShippingID:         shipping.ID,
		ShippingItems: 		shippingItems,
	}
	res := a.db.Create(&shippingModel)
	if res.Error == nil {
		shipping.ID = int64(shippingModel.ID)
	}
	return res.Error
}
