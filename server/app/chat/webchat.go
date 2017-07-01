package main

import (
	"encoding/json"
	"log"
	"net/http"

	//"github.com/igm/vendor"
	"./vendor"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
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
	chat.Publish("[info] new participant joined chat")
	defer chat.Publish("[info] participant left chat")

	var closedSession = make(chan struct{})

	go func() {
		reader, _ := chat.SubChannel(nil, session.ID())
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(msg.(string)); err != nil {
					return
				}
			}
		}
	}()

	for {
		if msg, err := session.Recv(); err == nil {
			var t interface{}
			if err := json.Unmarshal([]byte(msg), &t); err != nil {
				panic(err)
			}
			chat.Publish(msg)
			continue
		}
		break
	}

	close(closedSession)
	log.Println("sockjs session closed")

}
