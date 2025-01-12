package services

import (
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/jmoiron/sqlx"
)

type ProductService struct {
	DB *sqlx.DB
}

func NewProductService(db *sqlx.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) Create(in *models.ProductIn) (*models.Product, error) {
	sql := `
select id, title, code
from public.categories
where id = $1`
	cat := models.Category{}
	err := s.DB.Get(&cat, sql, in.Category)
	if err != nil {
		return nil, err
	}

	sql = `
insert into public.products(title, category, price)
values($1, $2, $3)
returning id, title, price`
	out := models.Product{}
	out.Category = cat
	err = s.DB.Get(&out, sql, in.Title, in.Category, in.Price)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ProductService) GetAll(category string) ([]*models.Product, error) {
	sql := `
select 
	p.id, p.title, p.price, 
	c.id as category_id, c.title as category_title, c.code as category_code
from public.products as p
inner join public.categories as c on c.id = p.category
where coalesce($1, '') = '' or p.category = $1::uuid`
	products := []*models.Product{}
	rows, err := s.DB.Queryx(sql, category)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.ID, &product.Title, &product.Price,
			&product.Category.ID, &product.Category.Title, &product.Category.Code)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (s *ProductService) GetByID(id string) (*models.Product, error) {
	sql := `
select id, title, category, price
from public.products
where id = $1::uuid`
	row, err := s.DB.Query(sql, id)
	if err != nil {
		return nil, err
	}
	product := models.Product{}
	if row.Next() {
		var categoryID string
		err = row.Scan(&product.ID, &product.Title, &categoryID, &product.Price)
		if err != nil {
			return nil, err
		}
		sql = `
select id, title, code
from public.categories
where id = $1::uuid`
		row, err = s.DB.Query(sql, categoryID)
		if err != nil {
			return nil, err
		}
		if row.Next() {
			category := models.Category{}
			err = row.Scan(&category.ID, &category.Title, &category.Code)
			if err != nil {
				return nil, err
			}
			product.Category = category
		}
	}
	return &product, nil
}

func (s *ProductService) Delete(id string) error {
	sql := `
delete from public.products
where id = $1::uuid`
	_, err := s.DB.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
