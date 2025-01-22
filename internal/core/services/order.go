package services

import (
	"encoding/json"
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jmoiron/sqlx"
	"log"
)

type OrderService struct {
	DB       *sqlx.DB
	Consumer *kafka.Consumer
	Topic    string
}

func NewOrderService(db *sqlx.DB, kafkaHost string, topic string) *OrderService {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaHost,
		"group.id":          "product-service",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create kafka consumer: %s", err)
	}
	return &OrderService{
		DB:       db,
		Consumer: c,
		Topic:    topic,
	}
}

func (o *OrderService) ShutDown() {
	if err := o.Consumer.Close(); err != nil {
		log.Fatalf("Failed to close kafka consumer: %s", err)
	}
}

func (o *OrderService) Run() {
	if err := o.Consumer.SubscribeTopics([]string{o.Topic}, nil); err != nil {
		log.Fatalf("Failed subscribe to kafka topics: %s", err)
	}

	for {
		msg, err := o.Consumer.ReadMessage(-1)
		if err != nil {
			log.Fatalf("Failed to read message in kafka consumer: %s", err)
		}

		order := models.OrderIn{}
		if err = json.Unmarshal(msg.Value, &order); err != nil {
			log.Fatalf("Failed to unmarshalling order: %s", err)
		}

		sql := `
insert into public.orders(product_id, user_id, amount)
values($1, $2, $3)
returing id, product_id, user_id, amount`
		_, err = o.DB.Query(sql, order.ProductID, order.UserID, order.Amount)
		if err != nil {
			log.Fatalf("Failed to make order: %s", err)
		}
	}
}
