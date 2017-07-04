package actions

import (
	"../vendor"
)

type PipeData struct {
	Msg     vendor.MsgSchema
	Session string
}

var ActionsMap map[string]func(PipeData) vendor.MsgSchema = map[string]func(PipeData) vendor.MsgSchema{

	"get-conversations": GetConversations,
	"chat-leave":        ChatLeave,
	"chat-join":         ChatJoin,
	"chat-open":         ChatOpen,
}

func GetConversations(pipe PipeData) vendor.MsgSchema {
	pipe.Msg["conversations"] = vendor.ChatApp.Conversations
	return pipe.Msg
}

//var ChatApp PipeData
