package services

import (
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/jmoiron/sqlx"
)

type CategoryService struct {
	DB *sqlx.DB
}

func NewCategoryService(db *sqlx.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) Create(in *models.CategoryIn) (*models.Category, error) {
	sql := `
insert into public.categories(title, code)
values($1, $2)
returning id, title, code`
	category := models.Category{}
	if err := s.DB.Get(&category, sql, in.Title, in.Code); err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *CategoryService) GetAll() ([]*models.Category, error) {
	sql := `
select id, title, code
from public.categories`
	categories := make([]*models.Category, 0)
	if err := s.DB.Select(&categories, sql); err != nil {
		return nil, err
	}
	return categories, nil
}
