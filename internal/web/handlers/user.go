package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/core/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		u.create(w, r)
	case http.MethodGet:
		u.getAll(w, r)
	}
}

func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	in := models.UserIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		msg := fmt.Sprintf("failed to unmarshalling user: %s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	out, err := u.UserService.Create(&in)
	if err != nil {
		msg := fmt.Sprintf("failed to create user: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&out); err != nil {
		msg := fmt.Sprintf("failed to marshalling user: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (u *UserHandler) getAll(w http.ResponseWriter, r *http.Request) {
	pageParams := models.NewPageParams(r.URL.RawQuery)
	users, err := u.UserService.GetAll(pageParams)
	if err != nil {
		msg := fmt.Sprintf("failed to get users: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(users); err != nil {
		msg := fmt.Sprintf("failed to marshalling users: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

}
