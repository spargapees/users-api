package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/spargapees/users-api/dto"
	"log"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	if db == nil {
		log.Print("Database connection is nil")
	}
	return &repository{db: db}
}

type Repository interface {
	CreateUser(user dto.User) (int, error)
}
