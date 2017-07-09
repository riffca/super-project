package actions

import (
	"../vendor"
)

type EmbedData struct {
	Msg     vendor.MsgSchema
	Session string
}

var ActionsMap map[string]func(EmbedData) vendor.MsgSchema = map[string]func(EmbedData) vendor.MsgSchema{
	"get-conversations": GetConversations,
	"chat-leave":        ChatLeave,
	"chat-join":         ChatJoin,
	"chat-open":         ChatOpen,
	"get-actions":       GetActions,
}

func GetConversations(data EmbedData) vendor.MsgSchema {
	data.Msg["conversations"] = vendor.ChatApp.Conversations
	return data.Msg
}

func GetActions(data EmbedData) vendor.MsgSchema {
	var m []string

	data.Msg["actions"] = m
	return data.Msg
}
