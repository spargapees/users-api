package service

import (
	"github.com/spargapees/users-api/dto"
	"github.com/spargapees/users-api/repository"
)

const (
	salt = "jafl85ytlashfl$8ghjlfd"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {

	return &service{repository: repository}
}

type Service interface {
	CreateUser(user dto.User) (int, error)
}
