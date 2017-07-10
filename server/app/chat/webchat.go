package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"github.com/igm/vendor"
	"./actions"
	"./vendor"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"strings"
)

var chat vendor.Publisher

func main() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(session sockjs.Session) {

	session.Send(clientConnect(session.ID()))
	chat.Publish(clientJoin(true))

	defer func() {
		vendor.ChatApp.RemoveConversation(session.ID())
		chat.Publish(clientJoin(false))
	}()

	var closedSession = make(chan struct{})

	go func() {
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
			var t vendor.MsgSchema
			if err := json.Unmarshal([]byte(msg), &t); err != nil {
				panic(err)
			}
			var p actions.EmbedData = actions.EmbedData{
				Msg:     t.Payload,
				Session: session.ID(),
			}
			t.Payload = actions.ActionsMap[t.Action](p)
			r, _ := json.Marshal(t)
			if strings.Index(t.Action, "chat") != -1 {
				chat.Publish(t)
			} else {
				session.Send(string(r))
			}
			continue
		}
		break
	}

	close(closedSession)
	log.Println("sockjs session closed")

}

//ROOT ACTIONS

func clientJoin(first bool) vendor.MsgSchema {
	m := make(map[string]interface{})
	var act string = "client-left"
	m["text"] = "[info] new participant left chat"
	if first {
		act = "client-join"
		m["text"] = "[info] new participant joined chat"
	}
	m["conversations"] = vendor.ChatApp.Conversations
	return vendor.MsgSchema{
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
	var fcm vendor.DataSchema = vendor.DataSchema{
		Action:  "client-connect",
		Payload: payload,
		Token:   "default",
	}
	result, _ := json.Marshal(fcm)
	return string(result)
}
