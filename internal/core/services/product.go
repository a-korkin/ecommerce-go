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
returning id, title, price;`
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
