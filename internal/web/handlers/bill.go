package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/models"
	pb "github.com/a-korkin/ecommerce/internal/proto"
	"github.com/a-korkin/ecommerce/internal/utils"
	"github.com/gofrs/uuid"
)

type BillHandler struct {
	Client *pb.BillServiceClient
}

func NewBillHandler(client *pb.BillServiceClient) *BillHandler {
	return &BillHandler{Client: client}
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

	in := pb.UserID{Id: userID}
	client := *b.Client
	bill, err := client.CreateBill(r.Context(), &in)
	if err != nil {
		msg := fmt.Sprintf("Failed to get bill by user: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	out := models.Bill{
		ID:       uuid.FromStringOrNil(bill.Id),
		TotalSum: bill.TotalPrice,
		Orders:   nil,
	}
	if err = json.NewEncoder(w).Encode(&out); err != nil {
		msg := fmt.Sprintf("Failed to marshalled bill: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
