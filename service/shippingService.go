package service

import (
	"encoding/json"
	d "main/domain"
	"main/utils/errs"
)

type ShippingService interface {
	PostShippingFee(body []byte) (*d.ShippingFee, *errs.AppError)
}

type DefaultShippingService struct {
	repo d.ShippingRepository
}

func NewShippingService(repo d.ShippingRepository) DefaultShippingService {
	return DefaultShippingService{repo}
}

func (s DefaultShippingService) PostShippingFee(body []byte) (*d.ShippingFee, *errs.AppError) {
	elem, _ := s.repo.GetShipping()
	if err := json.Unmarshal(body, elem); err != nil {
		return nil, errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	status := s.repo.ReturnSizeCost()
	if status != 200 {
		return nil, errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	status = s.repo.ReturnOptionCost()
	if status != 200 {
		return nil, errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	status = s.repo.ReturnRegionCost()
	if status != 200 {
		return nil, errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	_, fee := s.repo.GetShipping()
	return fee, nil
}
