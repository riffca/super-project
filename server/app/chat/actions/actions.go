package actions

import (
	"../core"
	//"fmt"
	decoder "github.com/mitchellh/mapstructure"
)

type Map interface{}

type EmbedData struct {
	Msg     Map
	Session string
	Token   string
}

var ActionsMap map[string]func(EmbedData) Map = map[string]func(EmbedData) Map{
	"get-conversations": GetConversations,
	// "chat-join":         ChatJoin,
	// "chat-send":         ChatSend,
	// "chat-create":       ChatCreate,
	"set-user":    SetUser,
	"get-metrics": GetMetrics,
	"user-create": UserCreate,
}

func UserCreate(data EmbedData) Map {
	type Payload struct {
		Name string `json:"name"`
	}
	p := Payload{}
	decoder.Decode(data.Msg, &p)
	return &p
}

func GetConversations(data EmbedData) Map {
	m := make(map[string]interface{})
	m["conversations"] = core.ChatApp.Conversations
	return m
}

func GetMetrics(data EmbedData) Map {
	m := make(map[string]interface{})
	m["conversations"] = core.ChatApp.Conversations
	m["users"] = core.Users
	m["connections"] = core.Connections
	return m
}

func SetUser(data EmbedData) Map {
	if data.Token == "default" {
		u := core.User{1, "default", "admin", data.Session}
		core.AddUser(u)
		return u
	}
	return new(interface{})
}
