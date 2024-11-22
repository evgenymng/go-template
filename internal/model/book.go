package model

import "github.com/google/uuid"

type Book struct {
	Id     uuid.UUID `json:"id"     db:"id"`
	Name   string    `json:"name"   db:"name"`
	Author string    `json:"author" db:"name"`
}
