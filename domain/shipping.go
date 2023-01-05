package domain

type ShippingFeeElem struct {
	Size   string `json:"size" validate:"required"`
	Option string `json:"option" validate:"required"`
	Region string `json:"region" validate:"required"`
}

type ShippingFee struct {
	ShippingFee uint `json:"shipping_fee"`
}

type ShippingRepository interface {
	GetShipping() (*ShippingFeeElem, *ShippingFee)
	ReturnSizeCost() int
	ReturnOptionCost() int
	ReturnRegionCost() int
}
