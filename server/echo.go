package main

import (
	schema "./schema"
	service "./service"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	//gorm "github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"

	pb "./service/chat/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
	"reflect"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var (
	websocket = flag.Bool("websocket", true, "enable/disable websocket protocol")
)

func init() {
	flag.Parse()
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

func main() {

	log.Println("Database connected:", schema.Connected)

	//contactChatService()
	opts := sockjs.DefaultOptions
	opts.Websocket = *websocket
	handler := sockjs.NewHandler("/echo", opts, echoHandler)
	http.Handle("/echo/", handler)
	http.HandleFunc("/admin/table", refreshTables)
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func refreshTables(w http.ResponseWriter, r *http.Request) {

	//------------------CODE EXAMPLE---------------------------------
	//https://astaxie.gitbooks.io/build-web-application-with-golang/en/03.2.html

	log.Println(`
    ---------------------------------------->
    ------------Http Handle-------------
    <--------------------------------------->`)

	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "<h1>ВСЕ ОК</h1>") // send data to client side

}

func echoHandler(session sockjs.Session) {

	log.Println("new sockjs session established")

	m := getMap()
	session.Send(string(m))

	for {

		if msg, err := session.Recv(); err == nil {

			var t DataSheme

			json.Unmarshal([]byte(msg), &t)

			log.Println("Request Data: ", t.RequestData)

			if val := service.CheckMethod(t.Service, t.Method); val == true {

				if t.Service == "Auth" {
					s := service.Auth{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})

				}

				if t.Service == "User" {
					s := service.User{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
				}

				if t.Service == "Page" {
					s := service.Page{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
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

// func activateAction(serviceName string, method string) {

//  fmt.Println(&service.Service)

//  r := reflect.ValueOf(&service.Service)

//  f := reflect.Indirect(r).FieldByName(serviceName)

//  reflect.ValueOf(&f).MethodByName(method).Call([]reflect.Value{})

// }

func contactChatService() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func getMap() {

	i := make(map[string]interface{})

	log.Println(service.MethodMap)

	for k, v := range service.MethodMap {
		i[k] = v
	}

	d := DataSheme{
		Service:      "All",
		Method:       "Methods",
		ResponseData: i,
	}

	r, _ := json.Marshal(&d)

	return r

}
