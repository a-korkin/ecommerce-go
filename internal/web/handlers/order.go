package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/gofrs/uuid"
	"net/http"
)

type OrderHandler struct{}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (o *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		o.create(w, r)
	}
}

func (o *OrderHandler) create(w http.ResponseWriter, r *http.Request) {
	in := models.OrderIn{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		msg := fmt.Sprintf("Failed to unmarshalled order: %s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Failed to create uuid: %v", http.StatusInternalServerError)
		return
	}
	out := models.Order{
		ID:        id,
		ProductID: in.ProductID,
		UserID:    in.UserID,
		Amount:    in.Amount,
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&out); err != nil {
		msg := fmt.Sprintf("Failed to marshalled order: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
