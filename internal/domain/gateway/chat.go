package gateway

import (
	"context"

	"github.com/nicholasferrer/chat-service/internal/domain/entities"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entities.Chat) error
	FindChatByID(ctx context.Context, chatID string) (*entities.Chat, error)
	SaveChat(ctx context.Context, chat *entities.Chat) error
}
