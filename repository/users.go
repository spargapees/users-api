package repository

import (
	"fmt"
	"github.com/spargapees/users-api/dto"
)

func (r *repository) CreateUser(user dto.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, email, password_hash) values ($1, $2, $3, $4) RETURNING id", "users")
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) GetAllUsers() ([]dto.User, error) {
	var users []dto.User
	query := fmt.Sprintf("SELECT id, name, email, password_hash FROM %s", "users")
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
