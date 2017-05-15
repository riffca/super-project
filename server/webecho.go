package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
	"reflect"
)

var (
	websocket = flag.Bool("websocket", true, "enable/disable websocket protocol")
)

func init() {
	flag.Parse()
}

func main() {
	opts := sockjs.DefaultOptions
	opts.Websocket = *websocket
	handler := sockjs.NewHandler("/echo", opts, echoHandler)
	http.Handle("/echo/", handler)
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

type DataSheme struct {
	Service string `json:"service"`
	Method  string `json:"method"`
	Token   string `json:"token"`
	MapId   string `json:"map_id"`
	Data    string `json:"data"`
}

func (t *DataSheme) CheckToken() {

	t.Token = "Token"

}

func echoHandler(session sockjs.Session) {

	log.Println("new sockjs session established")

	for {

		if msg, err := session.Recv(); err == nil {

			var t DataSheme

			json.Unmarshal([]byte(msg), &t)

			if err != nil {
				reflect.ValueOf(&t).MethodByName(t.Method).Call([]reflect.Value{})
			}
			response, _ := json.Marshal(&t)

			session.Send(string(response))

			continue

		}

		break

	}

	log.Println("sockjs session closed")

}
