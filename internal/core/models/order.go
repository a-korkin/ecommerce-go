package models

import "github.com/gofrs/uuid"

type Order struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID `json:"user_id"`
	Amount    int       `json:"amount"`
}

type OrderIn struct {
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID `json:"user_id"`
	Amount    int       `json:"amount"`
}
