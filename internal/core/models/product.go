package models

import "github.com/gofrs/uuid"

type Product struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Category Category  `json:"category"`
	Price    float32   `json:"price"`
}

type ProductIn struct {
	Title      string    `json:"title"`
	CategoryID uuid.UUID `json:"category_id"`
	Price      float32   `json:"price"`
}
