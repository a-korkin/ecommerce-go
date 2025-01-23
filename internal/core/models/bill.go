package models

import "github.com/gofrs/uuid"

type Bill struct {
	ID       uuid.UUID `json:"id"`
	TotalSum float32   `json:"total_sum"`
	Orders   []*Order  `json:"orders"`
}
