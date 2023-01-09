package service

import (
	"encoding/json"
	d "main/domain"
	"main/utils/auth"
	"main/utils/errs"
)

// primary port
type UserService interface {
	GetUsers() (*d.Users, error)
	ValidUser(body []byte) (string, *errs.AppError)
}

type DefaultUserService struct {
	repo d.UserRepository
}

func NewUserService(repo d.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}

func (s DefaultUserService) GetUsers() (*d.Users, error) {
	users, err := s.repo.Users()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s DefaultUserService) ValidUser(body []byte) (string, *errs.AppError) {
	account := &d.Account{}
	if err := json.Unmarshal(body, account); err != nil {
		return "", errs.NewBadRequestError(errs.PropErr, errs.PropMess)
	}
	if account.Password != "test01" {
		return "", errs.NewUnauthorizedError(errs.Unauthorized, errs.UnauthorizedMesss)
	}
	tStr := auth.SetClaims(account.Name)
	return tStr, nil
}
