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

	session.Send("[ + ]new sockjs session established " + session.ID())

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
			var t vendor.MsgSchema
			if err := json.Unmarshal([]byte(msg), &t); err != nil {
				panic(err)
			}
			var p actions.PipeData = actions.PipeData{
				Msg:     t,
				Session: session.ID(),
			}
			if action, ok := t["action"]; ok {
				m := actions.ActionsMap[action.(string)](p)
				r, _ := json.Marshal(m)

				if strings.Index(action.(string), "chat") != -1 {
					chat.Publish(m)
				} else {
					session.Send(string(r))
				}
				continue
			}
		}
		break
	}

	close(closedSession)
	log.Println("sockjs session closed")

}
