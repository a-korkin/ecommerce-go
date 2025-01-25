package repo

import "github.com/a-korkin/ecommerce/internal/core/models"

type ProductRepo interface {
	Create(*models.ProductIn) (*models.Product, error)
	Update(string, *models.ProductIn) (*models.Product, error)
	GetAll(*models.PageParams, string) ([]*models.Product, error)
	GetByID(string) (*models.Product, error)
	Delete(string) error
}
