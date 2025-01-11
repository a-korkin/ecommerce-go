package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/gofrs/uuid"
)

func newUUID() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

var categories []models.Category = []models.Category{
	{
		ID:    newUUID(),
		Title: "category#1",
		Code:  "cat#1",
	},
	{
		ID:    newUUID(),
		Title: "category#2",
		Code:  "cat#2",
	},
	{
		ID:    newUUID(),
		Title: "category#3",
		Code:  "cat#3",
	},
}

var products []*models.Product = []*models.Product{
	{
		ID:       newUUID(),
		Title:    "product#1",
		Category: categories[0],
		Price:    732.12,
	},
	{
		ID:       newUUID(),
		Title:    "product#2",
		Category: categories[2],
		Price:    62.23,
	},
	{
		ID:       newUUID(),
		Title:    "product#3",
		Category: categories[1],
		Price:    52.73,
	},
	{
		ID:       newUUID(),
		Title:    "product#4",
		Category: categories[1],
		Price:    591.51,
	},
	{
		ID:       newUUID(),
		Title:    "product#5",
		Category: categories[0],
		Price:    51.03,
	},
}

type ProductHandler struct {
	ProdService services.ProductService
}

func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		p.create(w, r)
	case "GET":
		getAll(w, r)
	}
}

// func ProductsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		create(w, r)
// 	case "GET":
// 		getAll(w, r)
// 	}
// }

// func ProductHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		getByID(w, r)
// 	case "DELETE":
// 		delete(w, r)
// 	}
// }

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	in := models.ProductIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Printf("failed to unmarshalling product: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProdService.Create(&in)
	if err != nil {
		log.Printf("failed to create product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// product := models.Product{
	// 	ID:       newUUID(),
	// 	Title:    in.Title,
	// 	Category: categories[0],
	// 	Price:    in.Price,
	// }
	// products = append(products, &product)
	if err := json.NewEncoder(w).Encode(&product); err != nil {
		log.Printf("failed to marshalling product: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getAll(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(&products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pID, err := uuid.FromString(id)
	if err != nil {
		log.Printf("failed to parse uuid: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, p := range products {
		if pID == p.ID {
			if err := json.NewEncoder(w).Encode(p); err != nil {
				log.Printf("failed to marshalling product: %v", err)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pID, err := uuid.FromString(id)
	if err != nil {
		log.Printf("failed to parse uuid: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	size := max(len(products)-1, 0)
	prods := make([]*models.Product, size)
	for _, p := range products {
		if pID != p.ID {
			prods = append(prods, p)
		}
	}
	if err := json.NewEncoder(w).Encode(&prods); err != nil {
		log.Printf("failed to marshalling products: %v", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}