package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/spargapees/users-api/dto"
)

func (s *service) CreateUser(user dto.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repository.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *service) GetAllUsers() ([]dto.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) UpdateUser(userId int, input dto.UserUpdate) error {
	return s.repository.UpdateUser(userId, input)
}
