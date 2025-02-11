package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/ports/repo"
	"github.com/a-korkin/ecommerce/internal/utils"
)

type ProductHandler struct {
	Repo repo.ProductRepo
}

func NewProductHandler(repo repo.ProductRepo) *ProductHandler {
	return &ProductHandler{
		Repo: repo,
	}
}

func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.create(w, r)
	case http.MethodPut:
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if ok {
			p.update(w, r, id)
		}
	case http.MethodGet:
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if !ok {
			p.getAll(w, r)
			return
		}
		p.getByID(w, r, id)
	case http.MethodDelete:
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.delete(w, r, id)
	}
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	in := models.ProductIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Printf("failed to unmarshalling product: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.Repo.Create(&in)
	if err != nil {
		log.Printf("failed to create product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&product); err != nil {
		log.Printf("failed to marshalling product: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *ProductHandler) update(
	w http.ResponseWriter, r *http.Request, id string) {
	in := models.ProductIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		msg := fmt.Sprintf("failed to unmarshalling product: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	out, err := h.Repo.Update(id, &in)
	if err != nil {
		msg := fmt.Sprintf("failed to update product: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(out); err != nil {
		msg := fmt.Sprintf("failed to marshalling product: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) getAll(w http.ResponseWriter, r *http.Request) {
	params := utils.GetQueryParams(r.URL.RawQuery)
	category := params["category"]
	pageParams := models.NewPageParams(r.URL.RawQuery)
	prods, err := h.Repo.GetAll(pageParams, category)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(&prods); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) getByID(
	w http.ResponseWriter, _ *http.Request, id string) {
	product, err := h.Repo.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) delete(
	w http.ResponseWriter, _ *http.Request, id string) {
	if err := h.Repo.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
