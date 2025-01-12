package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/utils"
)

type CategoryHandler struct {
	CategoryService *services.CategoryService
}

func NewCategoryHanlder(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: service}
}

func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.create(w, r)
	case "PUT":
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if !ok {
			msg := fmt.Sprintf("failed to get id")
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		h.update(w, r, id)
	case "GET":
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if ok {
			h.getByID(w, r, id)
			return
		}
		h.getAll(w, r)
	case "DELETE":
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if ok {
			h.delete(w, r, id)
		}
	}
}

func (h *CategoryHandler) create(w http.ResponseWriter, r *http.Request) {
	in := models.CategoryIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		msg := fmt.Sprintf("failed to unmarshalling category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	out, err := h.CategoryService.Create(&in)
	if err != nil {
		msg := fmt.Sprintf("failed to create category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&out); err != nil {
		msg := fmt.Sprintf("failed to marshalling category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) update(
	w http.ResponseWriter, r *http.Request, id string) {
	in := models.CategoryIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		msg := fmt.Sprintf("failed to unmarshalling category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	out, err := h.CategoryService.Update(id, &in)
	if err != nil {
		msg := fmt.Sprintf("failed to update category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&out); err != nil {
		msg := fmt.Sprintf("failed to marshalling category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) getAll(w http.ResponseWriter, _ *http.Request) {
	categories, err := h.CategoryService.GetAll()
	if err != nil {
		msg := fmt.Sprintf("failed to get categories: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(categories); err != nil {
		msg := fmt.Sprintf("failed to marshalling categories: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) getByID(
	w http.ResponseWriter, _ *http.Request, id string) {
	category, err := h.CategoryService.GetByID(id)
	if err != nil {
		msg := fmt.Sprintf("failed to get category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&category); err != nil {
		msg := fmt.Sprintf("failed to marshalling category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) delete(
	w http.ResponseWriter, _ *http.Request, id string) {
	if err := h.CategoryService.Delete(id); err != nil {
		msg := fmt.Sprintf("failed to delete category: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
