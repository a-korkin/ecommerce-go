package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/utils"
)

type BillHandler struct {
	BillService *services.BillService
}

func NewBillHandler(service *services.BillService) *BillHandler {
	return &BillHandler{BillService: service}
}

func (b *BillHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		b.getBillByUser(w, r)
	}
}

func (b *BillHandler) getBillByUser(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(r.RequestURI, "/{user_id}")
	userID, ok := vars["user_id"]
	if !ok {
		msg := fmt.Sprintf("User id not presented in url")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	out, err := b.BillService.GetBillByUser(userID)
	if err != nil {
		msg := fmt.Sprintf("Failed to get bill by user: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&out); err != nil {
		msg := fmt.Sprintf("Failed to marshalled bill: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
