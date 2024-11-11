package repository

import (
	"database/sql"

	"github.com/mauFade/call-doc/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(d *sql.DB) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO books (id, name, email, password) VALUES (?, ?, ?, ?)"

	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
