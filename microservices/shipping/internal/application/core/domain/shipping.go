package domain

import "time"

type ShippingItem struct {
	ProductCode 	string 			`json:"product_code"`
	Quantity   		int32  			`json:"quantity"`
}

type Shipping struct {
	ID            	int64          	`json:"id"`
	OrderID       	int64          	`json:"order_id"`
	DeliverySpan  	int32          	`json:"delivery_span"`
	ShippingItems 	[]ShippingItem 	`json:"shipping_items"`
	CreatedAt     	int64         	`json:"created_at"`
}

func NewShipping(orderId int64, shipping_items []ShippingItem) Shipping {
	delivery_span := GetDeliveryDays(shipping_items)
	return Shipping{
		OrderID:       orderId,
		ShippingItems: shipping_items,
		DeliverySpan:  delivery_span,
		CreatedAt:     time.Now().Unix(),
	}
}

func GetDeliveryDays(shipping_items []ShippingItem) int32 {
	totalUnits := int32(0)
	for _, item := range shipping_items {
		totalUnits += item.Quantity
	}
	return 1 + (totalUnits-1)/5
}