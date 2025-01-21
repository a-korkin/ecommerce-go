package services

import (
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) Create(in *models.UserIn) (*models.User, error) {
	sql := `
insert into public.users(first_name, last_name)
values($1, $2)
returning id, last_name, first_name;`
	row, err := s.DB.Query(sql, in.LastName, in.FirstName)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		out := models.User{}
		if err := row.Scan(&out.ID, &out.LastName, &out.FirstName); err != nil {
			return nil, err
		}
		return &out, nil
	}
	return nil, nil
}

func (s *UserService) GetAll(pageParams *models.PageParams) ([]*models.User, error) {
	sql := `
select id, last_name, first_name
from public.users
offset $1::integer * $2::integer
limit $2::integer`
	users := make([]*models.User, 0)
	err := s.DB.Select(
		&users, sql, pageParams.Page-1, pageParams.Limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}
