package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/rpc"
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
	Bills         *BillHandler
}

func NewRouter(
	db *sqlx.DB, kafkaProducer *kafka.Producer, grpcHost string) *Router {
	products := services.NewProductService(db)
	categories := services.NewCategoryService(db)
	users := services.NewUserService(db)
	client, err := rpc.NewGRPCClient(grpcHost)
	if err != nil {
		log.Fatalf("Failed to create grpc client: %s", err)
	}
	return &Router{
		KafkaProducer: kafkaProducer,
		Products:      NewProductHandler(products),
		Categories:    NewCategoryHandler(categories),
		Users:         NewUserHandler(users),
		Orders:        NewOrderHandler(kafkaProducer),
		Bills:         NewBillHandler(&client),
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
	case "/bills":
		router.Bills.ServeHTTP(w, r)
	default:
		fmt.Fprintf(w, "hello from main router\n")
	}
}
