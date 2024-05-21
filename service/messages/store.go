package messages

import (
	"context"

	"github.com/JerryJeager/ghost-write-backend/config"
	"github.com/JerryJeager/ghost-write-backend/models/messages"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageStore interface {
	CreateMessage(ctx context.Context, userID uuid.UUID, message *messages.Message) error
	GetMessages(ctx context.Context, userID uuid.UUID) (*messages.Messages, error)
}

type MessageRepo struct {
	client *gorm.DB
}

func NewMessageRepo(client *gorm.DB) *MessageRepo {
	return &MessageRepo{client: client}
}

func (o *MessageRepo) CreateMessage(ctx context.Context, userID uuid.UUID, message *messages.Message) error {
	result := config.Session.WithContext(ctx).Where("user_id = ?", userID).Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *MessageRepo) GetMessages(ctx context.Context, userID uuid.UUID) (*messages.Messages, error) {
	var messagesArr messages.Messages
	result := config.Session.WithContext(ctx).Where("user_id = ?", userID).Model(messages.Message{}).Find(&messagesArr)

	if result.Error != nil {
		return nil, result.Error
	}

	return &messagesArr, nil
}
