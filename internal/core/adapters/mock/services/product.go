package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/utils"
	"github.com/gofrs/uuid"
)

type ProductsMockService struct {
	Data []*models.Product
}

func NewProductsMockService() *ProductsMockService {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working dir: %s", err)
	}
	filePath := filepath.Join(currentDir, "../../../test", "products.json")
	data := make([]*models.Product, 4)
	utils.UnmarshallingFromFile(filePath, &data)

	return &ProductsMockService{Data: data}
}

func (s *ProductsMockService) Create(
	in *models.ProductIn) (*models.Product, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	prod := models.Product{
		ID:    id,
		Title: in.Title,
		Price: in.Price,
	}
	s.Data = append(s.Data, &prod)
	return &prod, nil
}

func (s *ProductsMockService) Update(
	id string, in *models.ProductIn) (*models.Product, error) {
	var prod *models.Product
	for _, d := range s.Data {
		if d.ID == uuid.FromStringOrNil(id) {
			prod = d
			break
		}
	}
	if prod != nil {
		prod.Title = in.Title
		prod.Price = in.Price
		return prod, nil
	}
	return nil, fmt.Errorf("Failed to find product by id: %s", id)
}

func (s *ProductsMockService) GetAll(
	pageParams *models.PageParams, category string) ([]*models.Product, error) {
	return s.Data, nil
}

func (s *ProductsMockService) GetByID(id string) (*models.Product, error) {
	for _, out := range s.Data {
		if out.ID == uuid.FromStringOrNil(id) {
			return out, nil
		}
	}
	return nil, fmt.Errorf("Failed to find product by id: %s", id)
}

func (s *ProductsMockService) Delete(id string) error {
	return nil
}
