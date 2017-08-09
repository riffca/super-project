package actions

import (
//"../core"
)

func ChatJoin(data EmbedData) Map {
	// 	adr := data.Msg["adress"]
	// 	room := core.ChatApp.Conversations[adr.(string)]
	// 	var ex bool = false
	// 	for _, m := range room.Members {
	// 		if m == data.Session {
	// 			ex = true
	// 			data.Msg["error"] = "User is already in room"
	// 		}
	// 	}
	// 	if !ex {
	// 		room.Members[data.Session] = data.Session
	// 		core.ChatApp.Conversations[adr.(string)] = room
	// 	}
	// 	data.Msg["chat"] = room
	// 	data.Msg["conversations"] = core.ChatApp.Conversations
	// 	return data.Msg
	// }

	// func ChatSend(data EmbedData) Map {
	// 	data.Msg["conversations"] = core.ChatApp.Conversations
	return data.Msg
}

func ChatCreate(data EmbedData) Map {

	// type Room struct {
	// 	Name     string
	// 	Members  map[string]string
	// 	Messages []Message
	// }

	// r := Room{
	// 	Name: data.Msg["name"].(string),
	// }

	// core.ChatApp

	// data.Msg["name"].(string)
	// core.ChatApp.

	return data.Msg

}
