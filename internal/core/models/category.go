package models

import "github.com/gofrs/uuid"

type Category struct {
	ID    uuid.UUID
	Title string
	Code  string
}
