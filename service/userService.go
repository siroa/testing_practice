package service

import (
	d "main/domain"
)

// primary port
type UserService interface {
	GetUsers() (*d.Users, error)
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
