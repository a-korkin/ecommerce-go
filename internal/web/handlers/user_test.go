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
	id := "9d2914e4-685a-4a59-91cd-e5a5b1d52c28"
	data := []byte(`{"last_name":"upd_last_name", "first_name":"upd_first_name"}`)
	r := httptest.NewRequest(http.MethodPut,
		fmt.Sprintf("/users/%s", id), bytes.NewBuffer(data))
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
	id := "7bb0c72b-6922-4d7b-839b-840ad3360442"
	r := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
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
	id := "9cb84f01-4835-48ed-9160-d34e12168c31"
	r := httptest.NewRequest(
		http.MethodDelete, fmt.Sprintf("/users/%s", id), nil)
	router.Users.delete(w, r, id)
	if status := w.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusNoContent)
	}
}
