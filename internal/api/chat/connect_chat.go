package chat

//func (i *ImplementationChat) ConnectChat(req *chat_v1.ConnectChatRequest, stream chat_v1.ChatV1_ConnectChatServer) error {
//	i.mxChannels.RLock()
//	chatChan, ok := i.channels[req.ChatId]
//	i.mxChannels.RUnlock()
//
//	if !ok {
//		return status.Errorf(codes.NotFound, "Chat with id %s not found", req.ChatId)
//	}
//
//	i.mxChat.Lock()
//	if _, okChat := i.channels[req.ChatId]; !okChat {
//
//	}
//}
