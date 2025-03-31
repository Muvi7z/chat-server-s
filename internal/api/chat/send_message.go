package chat

import (
	"context"
	"github.com/Muvi7z/chat-auth-s/gen/api/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *ImplementationChat) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	i.mxChannels.RLock()
	chatChan, ok := i.channels[req.GetCharId()]
	i.mxChannels.RUnlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "chat %s not found", req.GetCharId())
	}

	chatChan <- req.GetMessage()

	return &emptypb.Empty{}, nil

}
