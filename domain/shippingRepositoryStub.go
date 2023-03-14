package domain

import "time"

type ShippingRepositoryStub struct {
	ShippingFeeElem ShippingFeeElem
	ShippingFee     ShippingFee
}

func NewShippingRepositoryStub() *ShippingRepositoryStub {
	return &ShippingRepositoryStub{
		ShippingFeeElem: ShippingFeeElem{},
		ShippingFee:     ShippingFee{},
	}
}

func (s *ShippingRepositoryStub) GetShipping() (*ShippingFeeElem, *ShippingFee) {
	return &s.ShippingFeeElem, &s.ShippingFee
}

func (s *ShippingRepositoryStub) ReturnSizeCost() int {
	if s.ShippingFeeElem.Size == "small" {
		s.ShippingFee.ShippingFee += 300
	} else if s.ShippingFeeElem.Size == "medium" {
		s.ShippingFee.ShippingFee += 600
	} else if s.ShippingFeeElem.Size == "large" {
		s.ShippingFee.ShippingFee += 800
	} else {
		return 400
	}
	return 200
}

func (s *ShippingRepositoryStub) ReturnOptionCost() int {
	if s.ShippingFeeElem.Option == "regular" {
		s.ShippingFee.ShippingFee += 0
	} else if s.ShippingFeeElem.Option == "designation" {
		s.ShippingFee.ShippingFee += 100
	} else if s.ShippingFeeElem.Option == "haste" {
		s.ShippingFee.ShippingFee += 300
	} else {
		return 400
	}
	return 200
}

func (s *ShippingRepositoryStub) ReturnRegionCost() int {
	// broken time over 47sec
	time.Sleep(time.Second * 47)
	region := s.ShippingFeeElem.Region
	if region == "tokyo" {
		s.ShippingFee.ShippingFee += 0
	} else if region == "kanto" || region == "toukai" || region == "kansai" {
		s.ShippingFee.ShippingFee += 500
	} else if region == "hokkaido" || region == "sikoku" || region == "kyusyu" {
		s.ShippingFee.ShippingFee += 1000
	} else if region == "touhoku" || region == "koushinetu" || region == "hokuriku" || region == "tyugoku" {
		s.ShippingFee.ShippingFee += 800
	} else if region == "okinawa" {
		s.ShippingFee.ShippingFee += 1200
	} else {
		return 400
	}
	return 200
}
