package main

import (
	"encoding/json"
	"flag"
	//"fmt"
	service "./service"
	shema "./shema"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	(*gorm.DB).Create(&shema.User{Name: "STAS"})

	opts := sockjs.DefaultOptions
	opts.Websocket = *websocket
	handler := sockjs.NewHandler("/echo", opts, echoHandler)
	http.Handle("/echo/", handler)
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

type Service struct {
	Auth service.Auth
	User service.User
}

type DataSheme struct {
	Service      string                 `json:"service"`
	Method       string                 `json:"method"`
	Token        string                 `json:"token"`
	MapId        string                 `json:"map_id"`
	RequestData  map[string]interface{} `json:"request_data"`
	ResponseData map[string]interface{} `json:"response_data"`
}

func (t *DataSheme) Echo() {

	t.ResponseData = t.RequestData
	t.RequestData = nil

}

func echoHandler(session sockjs.Session) {

	log.Println("new sockjs session established")

	for {

		if msg, err := session.Recv(); err == nil {

			var t DataSheme

			json.Unmarshal([]byte(msg), &t)

			log.Println(t.Service)

			if t.Service == "Auth" {
				s := service.Auth{"token"}
				if val := checkMethod(t.Service, t.Method); val == true {
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
				}
			}

			if t.Service == "User" {
				u := service.User{}
				if val := checkMethod(t.Service, t.Method); val == true {
					reflect.ValueOf(&u).MethodByName(t.Method).Call([]reflect.Value{})
				}
			}

			t.Echo()
			response, _ := json.Marshal(&t)
			session.Send(string(response))
			continue

		}

		break

	}

	log.Println("sockjs session closed")

}

func checkMethod(service string, name string) bool {

	all := map[string][]string{
		"User": {
			"Test",
		},
		"Auth": {
			"checkToken",
		},
	}

	val := false

	for _, s := range all[service] {
		if s == name {
			val = true
		}
	}
	return val
}
