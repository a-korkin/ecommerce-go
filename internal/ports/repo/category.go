package repo

import "github.com/a-korkin/ecommerce/internal/core/models"

type CategoryRepo interface {
	Create(*models.CategoryIn) (*models.Category, error)
	Update(string, *models.CategoryIn) (*models.Category, error)
	GetAll(*models.PageParams) ([]*models.Category, error)
	GetByID(string) (*models.Category, error)
	Delete(string) error
}
