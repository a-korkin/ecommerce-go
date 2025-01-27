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

type UsersMockService struct {
	Data []*models.User
}

func NewUsersMockService() *UsersMockService {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working dir: %s", err)
	}
	filePath := filepath.Join(currentDir, "../../../test", "users.json")
	data := make([]*models.User, 3)
	utils.UnmarshallingFromFile(filePath, &data)
	return &UsersMockService{Data: data}
}

func (s *UsersMockService) Create(in *models.UserIn) (*models.User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	user := models.User{
		ID:        id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}
	s.Data = append(s.Data, &user)
	return &user, nil
}

func (s *UsersMockService) GetAll(pageParams *models.PageParams) ([]*models.User, error) {
	return s.Data, nil
}

func (s *UsersMockService) Update(
	id string, in *models.UserIn) (*models.User, error) {
	for _, u := range s.Data {
		if u.ID == uuid.FromStringOrNil(id) {
			u.LastName = in.LastName
			u.FirstName = in.FirstName
			return u, nil
		}
	}
	return nil, fmt.Errorf("Failed to get user by id: %s", id)
}

func (s *UsersMockService) GetByID(id string) (*models.User, error) {
	for _, u := range s.Data {
		if u.ID == uuid.FromStringOrNil(id) {
			return u, nil
		}
	}
	return nil, fmt.Errorf("Failed to get use by id: %s", id)
}

func (s *UsersMockService) Delete(string) error {
	return nil
}
