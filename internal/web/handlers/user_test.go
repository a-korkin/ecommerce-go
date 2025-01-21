package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/models"
)

func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	data := []byte(`{"last_name":"Smith", "first_name":"John"}`)
	r := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(data))
	router.Users.create(w, r)
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code, want: %v, got: %v",
			http.StatusCreated, status)
	}
	out := models.User{}
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling user: %s", err)
	}
	if out.LastName != "Smith" || out.FirstName != "John" {
		t.Errorf("Wrong user returned: %v", out)
	}
}

func TestGetAllUsers(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/users", nil)
	router.Users.getAll(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, want: %v, got: %v",
			http.StatusOK, status)
	}
	users := make([]*models.User, 0)
	if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
		t.Errorf("Failed to unmarshalling users: %s", err)
	}
	if len(users) < 2 {
		t.Errorf("Wrong count of users returned")
	}
}

func TestUpdateUser(t *testing.T) {
	w := httptest.NewRecorder()
	data := []byte(`{"last_name":"upd_last_name", "first_name":"upd_first_name"}`)
	r := httptest.NewRequest(http.MethodPut,
		"/users/4636a25d-02ee-4eb8-9757-efd677677076", bytes.NewBuffer(data))
	id := "4636a25d-02ee-4eb8-9757-efd677677076"
	router.Users.update(w, r, id)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.User{}
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling user: %s", err)
	}
	if out.LastName != "upd_last_name" || out.FirstName != "upd_first_name" {
		t.Errorf("Wrong user returned: %v", out)
	}
}

func TestGetByIDUser(t *testing.T) {
	w := httptest.NewRecorder()
	id := "5e782875-4d9c-4641-be3c-afddeb05c083"
	r := httptest.NewRequest(
		http.MethodGet, "/users/5e782875-4d9c-4641-be3c-afddeb05c083", nil)
	router.Users.getByID(w, r, id)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, want: %v, got: %v",
			http.StatusOK, status)
	}
	out := models.User{}
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling user: %s", err)
	}
	if out.LastName != "Petrov" || out.FirstName != "Petr" {
		t.Errorf("Wrong user returned: %v", out)
	}
}

func TestDeleteUser(t *testing.T) {
	w := httptest.NewRecorder()
	id := "d3f729cb-43c0-40c4-9084-74fb2b0bd408"
	r := httptest.NewRequest(
		http.MethodDelete, "/users/d3f729cb-43c0-40c4-9084-74fb2b0bd408", nil)
	router.Users.delete(w, r, id)
	if status := w.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusNoContent)
	}
}
