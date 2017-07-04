package actions

import (
	"../vendor"
)

type ChatCore interface {
	Signup()
	Login()
	All()
	Create()
	Join()
	Send()
}

type C ChatCore

func ChatJoin(pipe PipeData) vendor.MsgSchema {
	adr := pipe.Msg["adress"]
	room := vendor.ChatApp.Conversations[adr.(string)]
	var ex bool = false
	for _, m := range room.Members {
		if m == pipe.Session {
			ex = true
			pipe.Msg["error"] = "User is already in room"
		}
	}
	if !ex {
		room.Members[pipe.Session] = pipe.Session
		vendor.ChatApp.Conversations[adr.(string)] = room
	}
	pipe.Msg["active"] = room
	return pipe.Msg
}

func ChatLeave(pipe PipeData) vendor.MsgSchema {
	return pipe.Msg
}

func ChatOpen(pipe PipeData) vendor.MsgSchema {
	return pipe.Msg
}
