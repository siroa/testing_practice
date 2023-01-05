package service

import (
	"encoding/json"
	d "main/domain"
	"main/utils/errs"
)

// primary port
type OrderService interface {
	PostOrders(body []byte) (*d.Receipt, *errs.AppError)
}

type DefaultOrderService struct {
	repo d.OrderRepository
}

func NewOrderService(repo d.OrderRepository) DefaultOrderService {
	return DefaultOrderService{repo}
}

func (s DefaultOrderService) PostOrders(body []byte) (*d.Receipt, *errs.AppError) {
	p, r := s.repo.GetOrder()
	if err := json.Unmarshal(body, p); err != nil {
		return nil, errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	if p.Price == 0 {
		return nil, errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	s.repo.SetID()
	status := s.repo.CalcOrder()
	if status == 404 {
		return nil, errs.NewNotFoundError(errs.NotFond, errs.NotFondMess)
	}
	return r, nil
}
