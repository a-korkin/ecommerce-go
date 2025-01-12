package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/utils"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Products *ProductHandler
}

func NewRouter(db *sqlx.DB) *Router {
	productService := services.NewProductService(db)
	return &Router{
		Products: NewProductHandler(productService),
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch utils.GetResouce(r.RequestURI) {
	case "/products":
		router.Products.ServeHTTP(w, r)
	default:
		fmt.Fprintf(w, "hello from main router\n")
	}
}
