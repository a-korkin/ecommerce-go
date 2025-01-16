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

func TestUpdateProduct(t *testing.T) {
	id := "fd3310fd-2101-445f-ad3d-216fda4bd8a2"
	w := httptest.NewRecorder()
	data := []byte(`
{
	"title":"upd title",
	"category":"996be659-81f0-457c-8682-800abcfd64c2",
	"price": 32.21
}`)
	r := httptest.NewRequest(
		http.MethodPut, fmt.Sprintf("/products/%s", id), bytes.NewBuffer(data))
	router.Products.update(w, r, id)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.Product{}
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Errorf("Failed unmarshalling product: %s", err)
	}
	if out.Title != "upd title" || out.Price != 32.21 {
		t.Errorf("Wrong unmarshalling product: %v", out)
	}
}

func TestDeleteProduct(t *testing.T) {
	id := "6022261d-c88f-4551-8419-7319eb3ce18f"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		http.MethodDelete, fmt.Sprintf("/products/%s", id), nil)
	router.Products.delete(w, r, id)
	if status := w.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusNoContent)
	}
}
