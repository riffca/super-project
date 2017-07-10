package actions

import (
	"../vendor"
)

type Map map[string]interface{}

type EmbedData struct {
	Msg     Map
	Session string
}

var ActionsMap map[string]func(EmbedData) Map = map[string]func(EmbedData) Map{
	"get-conversations": GetConversations,
	"chat-join":         ChatJoin,
	"chat-send":         ChatSend,
}

func GetConversations(data EmbedData) Map {
	data.Msg["conversations"] = vendor.ChatApp.Conversations
	return data.Msg
}
