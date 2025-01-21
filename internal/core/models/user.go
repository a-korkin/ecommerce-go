package models

import "github.com/gofrs/uuid"

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
}

type UserIn struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
