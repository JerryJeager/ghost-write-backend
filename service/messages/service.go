package messages

import (
	"context"

	"github.com/JerryJeager/ghost-write-backend/models/messages"
	"github.com/google/uuid"
)

type MessageSv interface {
	CreateMessage(ctx context.Context, userID uuid.UUID, message *messages.Message) error
	GetMessages(ctx context.Context, userID uuid.UUID) (*messages.Messages, error)
}

type MessageServ struct {
	repo MessageStore
}

func NewMessageService(repo MessageStore) *MessageServ {
	return &MessageServ{repo: repo}
}

func (o *MessageServ) CreateMessage(ctx context.Context, userID uuid.UUID, message *messages.Message) error {
	return o.repo.CreateMessage(ctx, userID, message)
}

func (o *MessageServ) GetMessages(ctx context.Context, userID uuid.UUID) (*messages.Messages, error) {
	return o.repo.GetMessages(ctx, userID)
}
