package models

import "github.com/gofrs/uuid"

type Product struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Category Category  `json:"category" db:"category"`
	Price    float32   `json:"price" db:"price"`
}

type ProductIn struct {
	Title    string    `json:"title" db:"title"`
	Category uuid.UUID `json:"category" db:"category"`
	Price    float32   `json:"price" db:"price"`
}
