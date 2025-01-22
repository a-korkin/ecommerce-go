package handlers

import (
	"net/http"
)

type OrderHandler struct{}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (o *OrderHandler) ServeHTTP(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("hello from order handler"))
}
