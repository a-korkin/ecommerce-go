package services

import (
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/models"
)

func TestCreateCategory(t *testing.T) {
	in := models.CategoryIn{
		Title: "category@1",
		Code:  "cat@1",
	}
	out, err := categoryService.Create(&in)
	if err != nil {
		t.Errorf("Failed create category: %s", err)
	}
	if out.Title != in.Title || out.Code != in.Code {
		t.Errorf("Wrong saving category: %v", out)
	}
}

func TestUpdateCategory(t *testing.T) {
	id := "688e64d3-c722-48e5-be96-850e419df2d6"
	in := models.CategoryIn{
		Title: "upd category@1",
		Code:  "upd cat@1",
	}
	out, err := categoryService.Update(id, &in)
	if err != nil {
		t.Errorf("Failed to update category: %v", err)
	}
	if out.Title != "upd category@1" || out.Code != "upd cat@1" {
		t.Errorf("Wrong updating category: %v", out)
	}
}

func TestGetAllCategory(t *testing.T) {
	pageParams := models.NewPageParams("")
	out, err := categoryService.GetAll(pageParams)
	if err != nil {
		t.Errorf("Failed to get all categories: %v", err)
	}
	if len(out) < 2 {
		t.Errorf("Count of categories wrong, got: %d", len(out))
	}
}

func TestGetByIDCategory(t *testing.T) {
	id := "996be659-81f0-457c-8682-800abcfd64c2"
	out, err := categoryService.GetByID(id)
	if err != nil {
		t.Errorf("Failed to get category: %v", err)
	}
	if out.Title != "category@2" || out.Code != "cat@2" {
		t.Errorf("Wrong getted category: %v", out)
	}
}

func TestDeleteCategory(t *testing.T) {
	id := "efa8b389-a3bd-4e06-84dd-4960a0dfc55b"
	if err := categoryService.Delete(id); err != nil {
		t.Errorf("Failed to delete category: %s", err)
	}
}
