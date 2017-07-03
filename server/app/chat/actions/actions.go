package actions

import (
	"../vendor"
)

type PipeData struct {
	Msg     vendor.MsgSchema
	Session string
}

var ActionsMap map[string]func(PipeData) vendor.MsgSchema = map[string]func(PipeData) vendor.MsgSchema{

	"get-conversations": func(pipe PipeData) vendor.MsgSchema {
		pipe.Msg["conversations"] = vendor.ChatApp.Conversations
		return pipe.Msg

	},

	"chat-join": func(pipe PipeData) vendor.MsgSchema {
		adr := pipe.Msg["adress"]
		room := vendor.ChatApp.Conversations[adr.(string)]
		room = append(room, pipe.Session)
		pipe.Msg["active"] = room
		return pipe.Msg
	},

	"chat-send": func(pipe PipeData) vendor.MsgSchema {
		//pipe.Msg.ChatID
		return pipe.Msg
	},
}
