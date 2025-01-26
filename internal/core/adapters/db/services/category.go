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

func (s *CategoryService) Update(
	id string, in *models.CategoryIn) (*models.Category, error) {
	sql := `
update public.categories
set title = $2,
	code = $3
where id = $1::uuid
returning id, title, code`
	out := models.Category{}
	if err := s.DB.Get(&out, sql, id, in.Title, in.Code); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *CategoryService) GetAll(pageParams *models.PageParams) ([]*models.Category, error) {
	sql := `
select id, title, code
from public.categories
offset $1::integer * $2::integer
limit $2::integer`
	categories := make([]*models.Category, 0)
	err := s.DB.Select(&categories, sql, pageParams.Page-1, pageParams.Limit)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) GetByID(id string) (*models.Category, error) {
	sql := `
select id, title, code
from public.categories
where id = $1::uuid`
	out := models.Category{}
	if err := s.DB.Get(&out, sql, id); err != nil {
		return nil, err
	}
	return &out, nil

}

func (s *CategoryService) Delete(id string) error {
	sql := `
delete from public.categories
where id = $1::uuid`
	_, err := s.DB.Exec(sql, id)
	return err
}
