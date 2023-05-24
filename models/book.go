package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID         int         `db:"id" json:"id"`
	Date       time.Time   `db:"date" json:"date"`
	UserID     uuid.UUID   `db:"user_id" json:"user_id"`
	Appartment *Appartment `db:"-" json:"appartment,omitempty"`
}
