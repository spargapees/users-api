package service

import "github.com/spargapees/users-api/repository"

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {

	return &service{repository: repository}
}

type Service interface {
}
