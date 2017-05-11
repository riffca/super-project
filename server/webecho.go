package main

import (
	//"encoding/json"
	"flag"
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
	Data_Type     string            `json:"data_type"`
	Action_String string            `json:"action_string"`
	Log_List      map[string]string `json:"log_list"`
	Request_Map   map[string]string `json:"request_map"`
	Trans_Map     map[string]string `json:"trans_map"`
}

func (t *DataSheme) Foo() {
	log.Println("t.Data_Type")
}

func echoHandler(session sockjs.Session) {
	var t DataSheme
	log.Println("new sockjs session established")
	for {
		if msg, err := session.Recv(); err == nil {

			reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})

			session.Send(msg)
			continue
		}
		break
	}
	log.Println("sockjs session closed")
}
