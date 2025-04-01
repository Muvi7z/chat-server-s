package chat

import (
	"github.com/Muvi7z/chat-auth-s/gen/api/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImplementationChat) ConnectChat(req *chat_v1.ConnectChatRequest, stream chat_v1.ChatV1_ConnectChatServer) error {
	i.mxChannels.RLock()
	chatChan, ok := i.channels[req.ChatId]
	i.mxChannels.RUnlock()

	if !ok {
		return status.Errorf(codes.NotFound, "Chat with id %s not found", req.ChatId)
	}

	i.mxChat.Lock()
	if _, okChat := i.chats[req.ChatId]; !okChat {
		i.chats[req.GetChatId()] = &Chat{
			streams: make(map[string]chat_v1.ChatV1_ConnectChatServer),
		}
	}
	i.mxChat.Unlock()

	i.chats[req.GetChatId()].m.Lock()
	i.chats[req.GetChatId()].streams[req.GetUsername()] = stream
	i.chats[req.GetChatId()].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			if !okCh {
				return nil
			}

			for _, st := range i.chats[req.GetChatId()].streams {
				if err := st.Send(msg); err != nil {
					return err
				}

			}
		case <-stream.Context().Done():
			i.chats[req.GetChatId()].m.Lock()
			delete(i.chats[req.GetChatId()].streams, req.GetUsername())
			i.chats[req.GetChatId()].m.Unlock()
			return nil
		}
	}

}
