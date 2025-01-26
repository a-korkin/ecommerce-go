package repo

import "github.com/a-korkin/ecommerce/internal/core/models"

type UserRepo interface {
	Create(*models.UserIn) (*models.User, error)
	GetAll(*models.PageParams) ([]*models.User, error)
	Update(string, *models.UserIn) (*models.User, error)
	GetByID(string) (*models.User, error)
	Delete(string) error
}
