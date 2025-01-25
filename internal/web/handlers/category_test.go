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

// [
//     {
//         "id": "5d759639-093c-4b6b-9f45-2e586539464e",
//         "title": "category#1",
//         "code": "cat#1"
//     },
//     {
//         "id": "29e2f9a5-dbbf-4128-ba81-a27a8838bd9b",
//         "title": "category#2",
//         "code": "cat#2"
//     },
//     {
//         "id": "9256cf37-f395-46fc-86e8-33e99ce2ec60",
//         "title": "category#3",
//         "code": "cat#3"
//     },
//     {
//         "id": "d5ab1b04-23e3-4568-bc51-4ddf58630ea2",
//         "title": "category#4",
//         "code": "cat#4"
//     }
// ]

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
	id := "5d759639-093c-4b6b-9f45-2e586539464e"
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
	if out.Title != "category#1" || out.Code != "cat#1" {
		t.Errorf("Wrong unmarshalling category: %v", out)
	}
}

func TestCreateCategory(t *testing.T) {
	rr := httptest.NewRecorder()
	categoryData := []byte(`{"title":"category#5", "code":"cat#5"}`)
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
	if out.Title != "category#5" || out.Code != "cat#5" {
		t.Errorf("Wrong unmarshalling category, got: %v", out)
	}
}

func TestUpdateCategory(t *testing.T) {
	id := "29e2f9a5-dbbf-4128-ba81-a27a8838bd9b"
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
	id := "d5ab1b04-23e3-4568-bc51-4ddf58630ea2"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/categories", nil)
	router.Categories.delete(rr, req, id)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusNoContent)
	}
}
