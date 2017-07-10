package actions

import (
	"../vendor"
)

func ChatJoin(data EmbedData) Map {
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
	data.Msg["chat"] = room
	data.Msg["conversations"] = vendor.ChatApp.Conversations
	return data.Msg
}

func ChatSend(data EmbedData) Map {
	data.Msg["conversations"] = vendor.ChatApp.Conversations
	return data.Msg
}
