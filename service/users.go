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
