package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/utils"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	KafkaProducer *kafka.Producer
	Products      *ProductHandler
	Categories    *CategoryHandler
	Users         *UserHandler
	Orders        *OrderHandler
}

func NewRouter(db *sqlx.DB, kafkaProducer *kafka.Producer) *Router {
	products := services.NewProductService(db)
	categories := services.NewCategoryService(db)
	users := services.NewUserService(db)
	return &Router{
		KafkaProducer: kafkaProducer,
		Products:      NewProductHandler(products),
		Categories:    NewCategoryHanlder(categories),
		Users:         NewUserHandler(users),
		Orders:        NewOrderHandler(kafkaProducer),
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch utils.GetResouce(r.RequestURI) {
	case "/products":
		router.Products.ServeHTTP(w, r)
	case "/categories":
		router.Categories.ServeHTTP(w, r)
	case "/users":
		router.Users.ServeHTTP(w, r)
	case "/orders":
		router.Orders.ServeHTTP(w, r)
	default:
		fmt.Fprintf(w, "hello from main router\n")
	}
}
