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

func ChatJoin(data EmbedData) vendor.MsgSchema {
	adr := data.Msg["adress"]
	room := vendor.ChatApp.Conversations[adr.(string)]
	var ex bool = false
	for _, m := range room.Members {
		if m == data.Session {
			ex = true
			data.Msg["error"] = "User is already in room"
		}
	}
	if !ex {
		room.Members[data.Session] = data.Session
		vendor.ChatApp.Conversations[adr.(string)] = room
	}
	data.Msg["active"] = room
	return data.Msg
}

func ChatLeave(data EmbedData) vendor.MsgSchema {
	return data.Msg
}

func ChatOpen(data EmbedData) vendor.MsgSchema {
	return data.Msg
}
