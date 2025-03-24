package repository

import (
	"fmt"
	"github.com/spargapees/users-api/dto"
	"strings"
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
	query := fmt.Sprintf("SELECT id, name, surname, email, password_hash FROM %s", "users")
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) UpdateUser(userId int, input dto.UserUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, input.Surname)
		argId++
	}

	if len(setValues) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setQuery := strings.Join(setValues, ", ")

	args = append(args, userId)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", "users", setQuery, argId)

	_, err := r.db.Exec(query, args...)
	return err
}
