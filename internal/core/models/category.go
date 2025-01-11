package models

import "github.com/gofrs/uuid"

type Category struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	Code  string    `json:"code" db:"code"`
}
