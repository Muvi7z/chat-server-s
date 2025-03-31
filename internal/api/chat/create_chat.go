package chat

import (
	"context"
	"github.com/Muvi7z/chat-auth-s/gen/api/chat_v1"
	"github.com/google/uuid"
)

func (i *ImplementationChat) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {
	chatID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	i.channels[chatID.String()] = make(chan *chat_v1.Message, 100)

	return &chat_v1.CreateChatResponse{
		ChatId: chatID.String(),
	}, nil
}
