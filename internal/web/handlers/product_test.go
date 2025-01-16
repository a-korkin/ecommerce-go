package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-korkin/ecommerce/internal/core/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	productData := []byte(`
{
	"title":"product@10", 
	"category":"688e64d3-c722-48e5-be96-850e419df2d6", 
	"price":772.32
}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		http.MethodPost, "/products", bytes.NewBuffer(productData))
	router.Products.create(w, r)
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusCreated)
	}
	out := models.Product{}
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling product: %s", err)
	}
	if out.Title != "product@10" {
		t.Errorf("Expected: %s, got: %s", "product@10", out.Title)
	}
}

// ('5c0d6b4f-2d94-4e91-b69f-78f3832a810d', 'product@1', '688e64d3-c722-48e5-be96-850e419df2d6', 712.62),
// ('fd3310fd-2101-445f-ad3d-216fda4bd8a2', 'product@2', '688e64d3-c722-48e5-be96-850e419df2d6', 86.21),
// ('7ba8e565-82d9-4918-977c-85e62bc32e2c', 'product@3', '688e64d3-c722-48e5-be96-850e419df2d6', 23.31),
// ('6d49ce02-5e08-4d95-9451-c40bb44966e1', 'product@4', '996be659-81f0-457c-8682-800abcfd64c2', 73.25),
// ('85169049-293c-43e0-a0d9-327eeab730d4', 'product@5', 'efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 66.50),
// ('6022261d-c88f-4551-8419-7319eb3ce18f', 'product@6', '996be659-81f0-457c-8682-800abcfd64c2', 51.51),
// ('c4b129bd-43cb-4922-85ba-210dbd120ac3', 'product@7', 'efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 12.07),
// ('66724666-13ed-47f7-b042-eab7694e7499', 'product@8', '996be659-81f0-457c-8682-800abcfd64c2', 37.88),
// ('74958436-3427-4608-94b8-854c5db62e97', 'product@9', 'efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 3.63);

// id := "efa8b389-a3bd-4e06-84dd-4960a0dfc55b"
func TestGetByIDProduct(t *testing.T) {
	id := "5c0d6b4f-2d94-4e91-b69f-78f3832a810d"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", id), nil)
	router.Products.getByID(w, r, id)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.Product{}
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling product: %s", err)
	}
	if out.Title != "product@1" {
		t.Errorf("Expected: %s, got: %s", "product@1", out.Title)
	}
}

func TestGetAllProducts(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/products", nil)
	router.Products.getAll(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	products := make([]*models.Product, 10)
	if err := json.NewDecoder(w.Body).Decode(&products); err != nil {
		t.Errorf("Failed to unmarshalling products: %s", err)
	}
	if len(products) == 0 {
		t.Errorf("Handler doesn't return any products")
	}
}
