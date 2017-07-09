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

type DataSchema struct {
	Action  string           `json:"action"`
	Payload vendor.MsgSchema `json:"payload"`
	Token   string           `json:"token"`
}

func main() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(session sockjs.Session) {

	//CONNECT NEW CLIENT ACTION
	payload := make(map[string]interface{})
	payload["socket_session"] = session.ID()
	a := make([]string, 0)
	for k := range actions.ActionsMap {
		a = append(a, k)
	}
	payload["actions"] = a
	var fcm DataSchema = DataSchema{
		Action:  "client-connect",
		Payload: payload,
		Token:   "default",
	}
	result, _ := json.Marshal(fcm)
	session.Send(string(result))

	chat.Publish(vendor.MsgSchema{"text": "[info] new participant joined chat"})

	defer func() {
		vendor.ChatApp.RemoveConversation(session.ID())
		chat.Publish(vendor.MsgSchema{"text": "[info] participant left chat"})
	}()

	var closedSession = make(chan struct{})

	go func() {
		reader, _ := chat.SubChannel(nil, session.ID())
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

			var t DataSchema
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
				chat.Publish(t.Payload)
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
