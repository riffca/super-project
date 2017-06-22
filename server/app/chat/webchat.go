package main

import (
	"log"
	"net/http"

	"github.com/igm/pubsub"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

var chat pubsub.Publisher

type Room struct {
	*pubsub.Publisher
}

type Client struct {
	room    *pubsub.Publisher
	session *sockjs.Session
}

type Service struct {
	clients map[string]*Client
	message chan string
}

var clients []string

func (s *Service) appendClient(session sockjs.Session) {
	var puslisher pubsub.Publisher
	c := Client{&puslisher, &session}
	s.clients[session.ID()] = &c
}

func main() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(session sockjs.Session) {

	session.Send("[ + ]new sockjs session established " + session.ID())
	chat.Publish("[info] new participant joined chat")

	clients = append(clients, session.ID())
	log.Println(clients)

	var closedSession = make(chan struct{})

	defer chat.Publish("[info] participant left chat")

	go func() {

		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(msg.(string)); err != nil {
					log.Println("Session send")
					log.Println(msg.(string))
					return
				}
			}
		}
	}()
	for {
		if msg, err := session.Recv(); err == nil {
			chat.Publish(msg)
			continue
		}
		break
	}
	close(closedSession)
	log.Println("sockjs session closed")
}
