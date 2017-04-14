package main

import (
	"flag"
	"github.com/rs/cors"
	"log"
	"net/http"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

var (
	websocket = flag.Bool("websocket", true, "enable/disable websocket protocol")
)

func init() {
	flag.Parse()
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.

	handlerCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000"},
		AllowCredentials: true,
	}).Handler(mux)

	//handlerCors := cors.Default().Handler(mux)

	opts := sockjs.DefaultOptions
	opts.Websocket = *websocket
	handler := sockjs.NewHandler("/echo", opts, echoHandler)
	http.Handle("/echo/", handler)
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", handlerCors))
}

func echoHandler(session sockjs.Session) {
	log.Println("new sockjs session established")
	for {
		if msg, err := session.Recv(); err == nil {
			session.Send(msg)
			continue
		}
		break
	}
	log.Println("sockjs session closed")
}
