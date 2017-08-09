package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"github.com/igm/core"
	"./actions"
	"./core"
	"./models"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"strings"
)

var chat core.Publisher

func main() {
	connectDatabase()
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(session sockjs.Session) {

	core.AddConnection(session.ID())
	session.Send(clientConnect(session.ID()))
	chat.Publish(clientJoin(true))

	defer func() {
		core.ChatApp.RemoveConversation(session.ID())
		chat.Publish(clientJoin(false))
	}()

	var closedSession = make(chan struct{})

	go func() {
		//add new connections
		reader, _ := chat.SubChannel("close-subchannel", session.ID())
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				r, _ := json.Marshal(&msg)
				fmt.Println(string(r))
				if err := session.Send(string(r)); err != nil {
					return
				}
			}
		}
	}()

	for {
		if msg, err := session.Recv(); err == nil {
			actionListen(msg, session)
			continue
		}
		break
	}

	close(closedSession)
	log.Println("sockjs session closed")

}

func actionListen(msg string, session sockjs.Session) {

	//Get socket message
	var t core.MsgSchema
	if err := json.Unmarshal([]byte(msg), &t); err != nil {
		panic(err)
	}

	//make action use message
	var p actions.EmbedData = actions.EmbedData{
		Msg:     t.Payload,
		Session: session.ID(),
		Token:   t.Token,
	}
	//activate action, modify data

	for k, _ := range actions.ActionsMap {
		if k == t.Action {
			t.Payload = actions.ActionsMap[t.Action](p)
		}
	}

	r, _ := json.Marshal(t)
	//exec chat.publish in case
	if strings.Index(t.Action, "chat") != -1 {
		chat.Publish(t)
	} else {
		session.Send(string(r))
	}

}

//start data

func connectDatabase() {
	models.InitDB("127.0.0.1", "postgres", "chatsample", "1234", "5432")
	models.New().AutoMigrate(&models.User{})
}

//ROOT ACTIONS

func clientJoin(first bool) core.MsgSchema {

	m := make(map[string]interface{})
	var act string = "client-left"
	m["text"] = "[info] new participant left chat"
	if first {
		act = "client-join"
		m["text"] = "[info] new participant joined chat"
	}
	m["conversations"] = core.ChatApp.Conversations
	m["connections"] = core.Connections
	return core.MsgSchema{
		Action:  act,
		Payload: m,
	}
}

func clientConnect(id string) string {
	payload := make(map[string]interface{})
	payload["socket_session"] = id
	a := make([]string, 0)
	for k := range actions.ActionsMap {
		a = append(a, k)
	}
	payload["actions"] = a
	var fcm core.DataSchema = core.DataSchema{
		Action:  "client-connect",
		Payload: payload,
		Token:   "default",
	}
	result, _ := json.Marshal(fcm)
	return string(result)
}
