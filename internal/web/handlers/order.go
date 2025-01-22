package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofrs/uuid"
)

type OrderHandler struct {
	KafkaProducer *kafka.Producer
}

func NewOrderHandler(kafkaProducer *kafka.Producer) *OrderHandler {
	return &OrderHandler{KafkaProducer: kafkaProducer}
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
	value, err := json.Marshal(&out)
	if err != nil {
		msg := fmt.Sprintf("Failed to marshalled order: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	topic := "orders-v1-topic"
	err = o.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
	if err != nil {
		msg := fmt.Sprintf("Failed produce kafka: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order have been placed to broker\n"))

	// consumer
	// c, err := kafka.NewConsumer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092",
	// 	"group.id":          "product-service",
	// 	"auto.offset.reset": "earliest",
	// })
	// if err != nil {
	// 	msg := fmt.Sprintf("Failed to create kafka consumer: %s", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }
	// defer func() {
	// 	if err := c.Close(); err != nil {
	// 		log.Fatalf("Failed to close kafka consumer: %s", err)
	// 	}
	// }()
	// if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
	// 	msg := fmt.Sprintf("Failed to subscribe topics: %s", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }
	// msg, err := c.ReadMessage(-1)
	// if err != nil {
	// 	msg := fmt.Sprintf("Failed to read message from topic: %s", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }
	// result := models.Order{}
	// if err = json.Unmarshal(msg.Value, &result); err != nil {
	// 	msg := fmt.Sprintf("Failed to unmarshalling order in consumer: %s", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }
	// if err = json.NewEncoder(w).Encode(&result); err != nil {
	// 	msg := fmt.Sprintf("Failed to marshalling order in consumer: %s", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }
}
