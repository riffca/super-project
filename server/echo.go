package main

import (
	schema "./schema"
	service "./service"
	"encoding/json"
	"flag"
	"fmt"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"strings"

	pb "./service/chat/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/satori/go.uuid"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
	"reflect"
)

const (
	chatHost    = "localhost:50051"
	defaultName = "world"
)

var (
	websocket = flag.Bool("websocket", true, "enable/disable websocket protocol")
)

func init() {
	flag.Parse()
}

var sessions []uuid.UUID

var DB *gorm.DB

func InitDB(db *gorm.DB) {
	DB = db
}

type Back struct {
	SessionId string `json:"session_id"`
	Auth      bool   `json:"auth"`
	Json      string `json:"json"`
}

type DataScheme struct {
	Service      string                 `json:"service"`
	Method       string                 `json:"method"`
	Token        string                 `json:"token"`
	MapId        string                 `json:"map_id"`
	RequestData  map[string]interface{} `json:"request_data"`
	ResponseData map[string]interface{} `json:"response_data"`
	Back         Back                   `json:"back"`
	DB           *gorm.DB
}

func (t *DataScheme) DumpTables() {
	log.Println("Dump Tables")
	DB.DropTable(
		&schema.Page{},
		&schema.User{},
		&schema.Lead{},
		&schema.Message{})
	DB.DropTable("user_leads")
	DB.CreateTable(
		&schema.Page{},
		&schema.User{},
		&schema.Lead{},
		&schema.Message{})
}

func (t *DataScheme) Echo() {
	//needConvert()
	t.ResponseData = t.RequestData
	t.RequestData = nil

}

func (t *DataScheme) AddSession() {
	fmt.Println(t.Back.SessionId)
	fmt.Println(len(t.Back.SessionId))
	if len(t.Back.SessionId) > 5 {

	} else {
		ui := uuid.NewV4()
		t.Back.SessionId = ui.String()
		fmt.Println("append: ", t.Back.SessionId)
		sessions = append(sessions, ui)
	}

	fmt.Println("all sessions: ", sessions)
}

func (t *DataScheme) Auth() bool {
	t.Back.Auth = false
	return t.Back.Auth
}

func main() {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	} else {
		log.Println("Database connected")
	}
	db.LogMode(true)
	defer db.Close()
	service.InitDB(db)
	InitDB(db)
	db.CreateTable(
		&schema.Page{},
		&schema.User{},
		&schema.Lead{},
		&schema.Message{})

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

func echoHandler(session sockjs.Session) {

	log.Println("new sockjs session established")

	//send list of services
	session.Send(getServiceMap())

	for {

		if msg, err := session.Recv(); err == nil {

			var t DataScheme

			if err := json.Unmarshal([]byte(msg), &t); err != nil {
				panic(err)
			}

			t.AddSession()

			t.Auth()

			log.Println("RequestData: ", t.RequestData)

			if val := service.CheckMethod(t.Service, t.Method); val {

				switch t.Service {

				case "Auth":
					s := service.Auth{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
					//modify request data
					t.RequestData = s.Data
				case "Page":
					s := service.Page{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
					t.RequestData = s.Data
				case "User":
					s := service.User{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
					t.RequestData = s.Data
				case "Lead":
					s := service.Lead{Data: t.RequestData}
					reflect.ValueOf(&s).MethodByName(t.Method).Call([]reflect.Value{})
					t.RequestData = s.Data
				case "Data":
					t.DumpTables()
				default:
					t.RequestData["ERROR"] = "No server handler!"

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

func contactChatService() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(chatHost, grpc.WithInsecure())
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

func getServiceMap() string {

	i := make(map[string]interface{})

	for k, v := range service.MethodMap {
		i[k] = v
	}

	d := DataScheme{
		Service:      "Get",
		Method:       "Services",
		ResponseData: i,
	}

	r, _ := json.Marshal(&d)

	return string(r)

}

//---------------------------------------------------
//------------------CODE EXAMPLE---------------------
//---------------------------------------------------
//https://astaxie.gitbooks.io/build-web-application-with-golang/en/03.2.html
func refreshTables(w http.ResponseWriter, r *http.Request) {

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

//---------------------------------------------------
//---------------------------------------------------
//---------------------------------------------------

func needConvert(s interface{}) map[string]interface{} {

	inter := make(map[string]interface{})

	v := reflect.ValueOf(&s).Elem()

	for i := 0; i < v.NumField(); i++ {
		inter[v.Type().Field(i).Name] = v.Field(i).Interface()
	}

	return inter

}

func lear() {

	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

}
