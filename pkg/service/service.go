package service

import "github.com/NIKFED/go-auth-users/pkg/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
