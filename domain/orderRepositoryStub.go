package domain

import (
	"math/rand"
	"time"
)

type OrderRepositoryStub struct {
	Purchase Purchase
	Receipt  Receipt
}

func NewOrderRepositoryStub() *OrderRepositoryStub {
	return &OrderRepositoryStub{
		Purchase: Purchase{},
		Receipt:  Receipt{},
	}
}
func (o *OrderRepositoryStub) GetOrder() (*Purchase, *Receipt) {
	return &o.Purchase, &o.Receipt
}

func (o *OrderRepositoryStub) SetID() {
	rand.Seed(time.Now().UnixNano())
	o.Receipt.OrderID = int64(rand.Intn(99) + 100)
}

func (o *OrderRepositoryStub) CalcOrder() int {
	o.Receipt.ShippingFee = 300 //配送料は一律300円とする
	if o.Purchase.UserID == 101 {
		o.Receipt.SumPrice = DiscountedPrice(int(o.Purchase.Price), 0.1, 0.07, 0.03)
		if int(o.Purchase.Price) >= 5000 && o.Purchase.Contained {
			o.Receipt.ShippingFee = 0
		}
		o.Receipt.SumPrice += o.Receipt.ShippingFee
	} else if o.Purchase.UserID == 201 {
		o.Receipt.SumPrice = DiscountedPrice(int(o.Purchase.Price), 0.05, 0.03, 0.0)
		if int(o.Purchase.Price) >= 10000 && o.Purchase.Contained {
			o.Receipt.ShippingFee = 0
		}
		o.Receipt.SumPrice += o.Receipt.ShippingFee
	} else if o.Purchase.UserID == 301 {
		o.Receipt.SumPrice = DiscountedPrice(int(o.Purchase.Price), 0.03, 0.0, 0.0)
		o.Receipt.SumPrice += o.Receipt.ShippingFee
	} else {
		return 404
	}
	return 200
}

func DiscountedPrice(price int, a, b, c float64) int {
	sumPrice := 0
	if price >= 10000 {
		sumPrice = price - int(float64(price)*a)
	} else if price >= 5000 {
		sumPrice = price - int(float64(price)*b)
	} else {
		sumPrice = price - int(float64(price)*c)
	}
	return sumPrice
}
