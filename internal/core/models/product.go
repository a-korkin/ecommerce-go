package models

import "github.com/gofrs/uuid"

type Product struct {
	ID       uuid.UUID
	Title    string
	Category Category
	Price    float32
}
