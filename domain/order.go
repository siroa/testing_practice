package domain

type Purchase struct {
	UserID    int64 `json:"user_id" validate:"required,max=999"`
	Price     uint  `json:"price" validate:"required,min=0,max=4294967295"`
	Contained *bool `json:"contained" validate:"required"`
}

type Receipt struct {
	OrderID     int   `json:"order_id"`
	ShippingFee int64 `json:"shipping_fee"`
	SumPrice    int64 `json:"sum_price"`
}

// secondary port
type OrderRepository interface {
	GetOrder() (*Purchase, *Receipt)
	CalcOrder() int
	SetID()
}
