package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/utils"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Products   *ProductHandler
	Categories *CategoryHandler
}

func NewRouter(db *sqlx.DB) *Router {
	products := services.NewProductService(db)
	categories := services.NewCategoryService(db)
	return &Router{
		Products:   NewProductHandler(products),
		Categories: NewCategoryHanlder(categories),
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch utils.GetResouce(r.RequestURI) {
	case "/products":
		router.Products.ServeHTTP(w, r)
	case "/categories":
		router.Categories.ServeHTTP(w, r)
	default:
		fmt.Fprintf(w, "hello from main router\n")
	}
}
