package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/models"
)

func TestGetAllCategories(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	router.Categories.getAll(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
}

func TestGetByIDCategory(t *testing.T) {
	id := "688e64d3-c722-48e5-be96-850e419df2d6"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf("/categories/%s", id), nil)
	router.Categories.getByID(rr, req, id)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.Category{}
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling category: %v", err)
	}
	if out.Title != "category@1" || out.Code != "cat@1" {
		t.Errorf("Wrong unmarshalling category: %v", out)
	}
}

func TestCreateCategory(t *testing.T) {
	rr := httptest.NewRecorder()
	categoryData := []byte(`{"title":"category@4", "code":"cat@4"}`)
	req := httptest.NewRequest(http.MethodPost, "/categories",
		bytes.NewBuffer(categoryData))
	req.Header.Set("Content-Type", "application/json")
	router.Categories.create(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusCreated)
	}

	out := models.Category{}
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling category: %s", err)
	}
	if out.Title != "category@4" || out.Code != "cat@4" {
		t.Errorf("Wrong unmarshalling category, got: %v", out)
	}
}

func TestUpdateCategory(t *testing.T) {
	id := "996be659-81f0-457c-8682-800abcfd64c2"
	rr := httptest.NewRecorder()
	categoryData := []byte(`{"title":"upd title", "code":"upd code"}`)
	req := httptest.NewRequest(http.MethodPut,
		fmt.Sprintf("/categories/%s", id), bytes.NewBuffer(categoryData))
	req.Header.Set("Content-Type", "application/json")
	router.Categories.update(rr, req, id)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.Category{}
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling category: %s", err)
	}
	if out.Title != "upd title" || out.Code != "upd code" {
		t.Errorf("Wrong unmarshalling category, got: %v", out)
	}
}

func TestDeleteCategory(t *testing.T) {
	id := "efa8b389-a3bd-4e06-84dd-4960a0dfc55b"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/categories", nil)
	router.Categories.delete(rr, req, id)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusNoContent)
	}
}
