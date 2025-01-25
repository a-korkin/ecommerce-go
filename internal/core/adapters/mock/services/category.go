package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"

	"github.com/a-korkin/ecommerce/internal/core/models"
)

type CategoryMockService struct {
	Data []*models.Category
}

func NewCategoryMockService() (*CategoryMockService, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	testJSON := filepath.Join(currentDir, "test", "categories.json")
	data := make([]*models.Category, 4)
	file, err := os.Open(testJSON)
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return &CategoryMockService{Data: data}, nil
}

func (s *CategoryMockService) Create(
	in *models.CategoryIn) (*models.Category, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	cat := models.Category{
		ID:    id,
		Title: in.Title,
		Code:  in.Code,
	}
	s.Data = append(s.Data, &cat)
	return &cat, nil
}

func (s *CategoryMockService) Update(
	id string, in *models.CategoryIn) (*models.Category, error) {
	var cat *models.Category
	for _, d := range s.Data {
		if d.ID == uuid.FromStringOrNil(id) {
			cat = d
			break
		}
	}
	if cat != nil {
		cat.Title = in.Title
		cat.Code = in.Code
		return cat, nil
	}
	return nil, fmt.Errorf("Failed to find category by id: %s", id)
}

func (s *CategoryMockService) GetAll(
	pageParams *models.PageParams) ([]*models.Category, error) {
	return s.Data, nil
}

func (s *CategoryMockService) GetByID(id string) (*models.Category, error) {
	for _, out := range s.Data {
		if out.ID == uuid.FromStringOrNil(id) {
			return out, nil
		}
	}
	return nil, fmt.Errorf("Failed to find category by id: %s", id)
}

func (s *CategoryMockService) Delete(id string) error {
	return nil
}
