package messages

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Message   string     `json:"message" gorm:"required"`
	UserID    uuid.UUID  `json:"user_id"`
}

type Messages []Message
