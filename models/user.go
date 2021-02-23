package models

import (
	"github.com/gofrs/uuid"
)

// User model struct
type User struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

func (User) TableName() string {
	return "users"
}
