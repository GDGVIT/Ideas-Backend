package models

import (
	"time"

	"github.com/google/uuid"
)

// User model.
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email     string    `gorm:"type:varchar(100);unique_index;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
