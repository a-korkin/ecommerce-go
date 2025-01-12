package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/utils"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		p.create(w, r)
	case "PUT":
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if ok {
			p.update(w, r, id)
		}
	case "GET":
		path := "/{id}"
		vars := utils.GetVars(r.RequestURI, path)
		id, ok := vars["id"]
		if !ok {
			p.getAll(w, r)
			return
		}
		p.getByID(w, r, id)
	case "DELETE":
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
	product, err := h.ProductService.Create(&in)
	if err != nil {
		log.Printf("failed to create product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(&product); err != nil {
		log.Printf("failed to marshalling product: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) update(
	w http.ResponseWriter, r *http.Request, id string) {
	in := models.ProductIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		msg := fmt.Sprintf("failed to unmarshalling product: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	out, err := h.ProductService.Update(id, &in)
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

	prods, err := h.ProductService.GetAll(category)
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
	product, err := h.ProductService.GetByID(id)
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
	if err := h.ProductService.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
