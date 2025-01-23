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
}
