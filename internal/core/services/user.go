package services

import "github.com/jmoiron/sqlx"

type UserService struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{DB: db}
}
