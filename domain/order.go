package domain

type Purchase struct {
	UserID    int64 `json:"user_id" validate:"required"`
	Price     uint  `json:"price" validate:"required"`
	Contained bool  `json:"contained" validate:"required"`
}

type Receipt struct {
	OrderID     int64 `json:"order_id"`
	ShippingFee int   `json:"shipping_fee"`
	SumPrice    int   `json:"sum_price"`
}

// secondary port
type OrderRepository interface {
	GetOrder() (*Purchase, *Receipt)
	CalcOrder() int
	SetID()
}
