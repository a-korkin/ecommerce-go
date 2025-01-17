package services

import (
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/gofrs/uuid"
)

func TestCreateProduct(t *testing.T) {
	in := models.ProductIn{
		Title:    "product@10",
		Category: uuid.FromStringOrNil("996be659-81f0-457c-8682-800abcfd64c2"),
		Price:    5321.21,
	}
	out, err := productService.Create(&in)
	if err != nil {
		t.Errorf("Failed to create product: %s", err)
	}
	if out.Title != "product@10" || out.Price != 5321.21 {
		t.Errorf("Returned wrong product: %v", out)
	}
}

func TestUpdateProduct(t *testing.T) {
	id := "7ba8e565-82d9-4918-977c-85e62bc32e2c"
	in := models.ProductIn{
		Title:    "upd product@3",
		Category: uuid.FromStringOrNil("688e64d3-c722-48e5-be96-850e419df2d6"),
		Price:    512.66,
	}
	out, err := productService.Update(id, &in)
	if err != nil {
		t.Errorf("Failed to update product: %s", err)
	}
	if out.Title != "upd product@3" || out.Price != 512.66 {
		t.Errorf("Returned wrong product: %v", out)
	}
}

func TestGetAllProducts(t *testing.T) {
	categoryID := "688e64d3-c722-48e5-be96-850e419df2d6"
	out, err := productService.GetAll(categoryID)
	if err != nil {
		t.Errorf("Failed to get products by categoryID: %s", err)
	}
	if len(out) != 3 {
		t.Errorf("Returned wrong count of products, got: %d, want: %d", len(out), 3)
	}
}

func TestGetByIDProduct(t *testing.T) {
	id := "6022261d-c88f-4551-8419-7319eb3ce18f"
	out, err := productService.GetByID(id)
	if err != nil {
		t.Errorf("Failed to get product by id: %s", err)
	}
	if out.Title != "product@6" || out.Price != 51.51 {
		t.Errorf("Returned wront produt: %v", out)
	}
}

func TestDeleteProduct(t *testing.T) {
	id := "66724666-13ed-47f7-b042-eab7694e7499"
	if err := productService.Delete(id); err != nil {
		t.Errorf("Failed to delete product: %s", err)
	}
}
