package chat

import (
	"github.com/Muvi7z/chat-auth-s/gen/api/chat_v1"
	"sync"
)

type Chat struct {
	streams map[string]chat_v1.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

type ImplementationChat struct {
	chat_v1.UnimplementedChatV1Server

	chats  map[string]*Chat
	mxChat sync.RWMutex

	channels   map[string]chan *chat_v1.Message
	mxChannels sync.RWMutex
}

func NewImplementationChat() *ImplementationChat {
	return &ImplementationChat{
		chats:    make(map[string]*Chat),
		channels: make(map[string]chan *chat_v1.Message),
	}
}
