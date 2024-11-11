package service

import "database/sql"

type UserService struct {
	db *sql.DB
}

func NewUserService(d *sql.DB) *UserService {
	return &UserService{
		db: d,
	}
}

func (s *UserService) CreateUser() {}
